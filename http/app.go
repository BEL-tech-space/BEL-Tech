package http

import (
	"beltech/service"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type App struct {
	service             *service.Service
	accountSessionStore sessions.Store
	agentSessionStore   sessions.Store
	collectors          []collector
}

type collector struct {
	AgentID uuid.UUID
	Email   string
	Role    string
}

func NewApp(cfg EnvConfig, db *pgxpool.Pool) *App {

	store := sessions.NewCookieStore([]byte(cfg.CookieSecret))
	store.MaxAge(86400 * 30) // 30 days
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = true

	gothic.Store = store

	goth.UseProviders(
		google.New(cfg.LoginProvider.Google.ClientKey, cfg.LoginProvider.Google.Secret, "/auth/google/callback"),
		// facebook.New(cfg.LoginProvider.Facebook.ClientKey, cfg.LoginProvider.Facebook.Secret, "http://localhost:8090/auth/facebook/callback"),
	)

	app := App{}
	app.accountSessionStore = sessions.NewCookieStore([]byte(GenerateRandomKey(cfg.SessionSecret)))
	app.agentSessionStore = sessions.NewCookieStore([]byte(GenerateRandomKey(cfg.SessionSecret)))
	app.service = service.NewService(db)

	return &app
}

func GenerateRandomKey(d string) string {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("failed to generate random key:", err)
		return d
	}
	return hex.EncodeToString(key)
}
