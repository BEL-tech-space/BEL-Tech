package http

import (
	"beltech/service"
	"context"
	"slices"

	"github.com/dgraph-io/ristretto/v2"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type App struct {
	config              *EnvConfig
	service             *service.Service
	accountSessionStore sessions.Store
	agentSessionStore   sessions.Store
	collectors          []collector
	// beltechWatermarkLogo []byte
	accountCache *ristretto.Cache[string, service.Account]
}

type collector struct {
	AgentID uuid.UUID
	Email   string
	Role    string
}
type contextKey string

const (
	contextKeyEmail   contextKey = "email"
	contextKeyAgentID contextKey = "agentID"
	contextKeyRole    contextKey = "role"
)

const productionEnv string = "production"

func NewApp(cfg *EnvConfig, db *pgxpool.Pool, rCache *ristretto.Cache[string, service.Account]) *App {

	store := sessions.NewCookieStore([]byte(cfg.CookieSecret))
	store.MaxAge(86_400 * 15) // 30 days
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = true

	var callbackURL string = "http://localhost:8090/account/auth/google/callback"
	if cfg.Build == productionEnv {
		callbackURL = "https://beltech.site/account/auth/google/callback"
	}

	gothic.Store = store
	goth.UseProviders(
		google.New(cfg.LoginProvider.Google.ClientKey, cfg.LoginProvider.Google.Secret, callbackURL),
		// facebook.New(cfg.LoginProvider.Facebook.ClientKey, cfg.LoginProvider.Facebook.Secret, "http://localhost:8090/auth/facebook/callback"),
	)

	app := App{}
	app.config = cfg

	app.accountCache = rCache
	app.accountSessionStore = sessions.NewCookieStore([]byte(GenerateRandomKey(cfg.SessionSecret, cfg.Build)))
	app.agentSessionStore = sessions.NewCookieStore([]byte(GenerateRandomKey(cfg.SessionSecret, cfg.Build)))
	app.service = service.NewService(db)

	notifs, err := app.service.GetAllNotifications(context.TODO())
	if err != nil {
		ActiveNotifications = make([]service.Notification, 0, 50)
	} else {
		ActiveNotifications = slices.Grow(notifs, 50)
	}

	if cfg.Build == productionEnv {
		app.collectors = append(app.collectors, collector{Email: devGothUser.Email})
	}

	return &app
}

func GenerateRandomKey(d, buildMode string) string {
	// if buildMode != "production" && (len([]byte(d)) == 32 || len([]byte(d)) == 64) {
	// 	return d
	// }

	// key := make([]byte, 64)
	// if _, err := rand.Read(key); err != nil {
	// 	return d
	// }

	// return hex.EncodeToString(key)
	return d
}
