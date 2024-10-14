package http

import (
	"beltech/service"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/patrickmn/go-cache"
)

type App struct {
	config               *EnvConfig
	service              *service.Service
	accountSessionStore  sessions.Store
	agentSessionStore    sessions.Store
	collectors           []collector
	beltechWatermarkLogo []byte
	cacheStorage         *cache.Cache
}

type collector struct {
	AgentID uuid.UUID
	Email   string
	Role    string
}
type contextKey string

const (
	contextKeyEmail      contextKey = "email"
	contextUserAvatarURL contextKey = "avatarURL"
	contextKeyAgentID    contextKey = "agentID"
	contextKeyUserID     contextKey = "userID"
	contextKeyRole       contextKey = "role"
)

func NewApp(cfg *EnvConfig, db *pgxpool.Pool) *App {

	store := sessions.NewCookieStore([]byte(cfg.CookieSecret))
	store.MaxAge(86_400 * 30) // 30 days
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = true

	var callbackURL string
	if cfg.Build == productionEnv {
		callbackURL = "https://beltech.site/account/auth/google/callback"
	} else {
		callbackURL = "http://localhost:8090/account/auth/google/callback"
	}

	gothic.Store = store
	goth.UseProviders(
		google.New(cfg.LoginProvider.Google.ClientKey, cfg.LoginProvider.Google.Secret, callbackURL),
		// facebook.New(cfg.LoginProvider.Facebook.ClientKey, cfg.LoginProvider.Facebook.Secret, "http://localhost:8090/auth/facebook/callback"),
	)

	app := App{}
	app.config = cfg
	app.cacheStorage = cache.New(24*time.Hour, 24*time.Hour)
	app.accountSessionStore = sessions.NewCookieStore([]byte(GenerateRandomKey(cfg.SessionSecret, cfg.Build)))
	app.agentSessionStore = sessions.NewCookieStore([]byte(GenerateRandomKey(cfg.SessionSecret, cfg.Build)))
	app.service = service.NewService(db)

	return &app
}

func GenerateRandomKey(d, buildMode string) string {
	if buildMode != "production" && (len([]byte(d)) == 32 || len([]byte(d)) == 64) {
		return d
	}
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("failed to generate random key:", err)
		return d
	}
	return hex.EncodeToString(key)
}
