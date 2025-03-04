package http

import (
	"beltech/service"
	"context"
	"slices"

	"github.com/dgraph-io/ristretto/v2"
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
	collectors          map[string]service.Collector // key is collector's email
	accountCache        *ristretto.Cache[string, service.Account]
	wasteLogCache       *ristretto.Cache[string, service.WasteLog]
}

type contextKey string

const (
	contextKeyEmail   contextKey = "email"
	contextKeyAgentID contextKey = "agentID"
	contextKeyRole    contextKey = "role"
)

const productionEnv string = "production"

func NewApp(
	cfg *EnvConfig,
	db *pgxpool.Pool,
	accountCache *ristretto.Cache[string, service.Account],
	wasteLogCache *ristretto.Cache[string, service.WasteLog]) *App {

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

	app.accountCache = accountCache
	app.wasteLogCache = wasteLogCache
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
		if err := app.setupProductionModeCollectors(); err != nil {
			panic(err)
		}
	} else {
		if err := app.setupdDevModeCollectors(); err != nil {
			panic(err)
		}
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

func (a *App) setupdDevModeCollectors() error {
	devCollector := service.AddNewAgentAsCollectorParams{
		CollectorEmail: devGothUser.Email,
		DisplayName:    devGothUser.NickName,
		AvatarUrl:      devGothUser.AvatarURL,
	}

	collector, err := a.service.AddNewAgentAsCollector(
		context.Background(),
		devCollector,
	)
	if err != nil {
		return err
	}

	a.collectors = make(map[string]service.Collector, 2)

	a.collectors[devGothUser.Email] = service.Collector{
		AgentID:        collector.AgentID,
		CollectorEmail: devGothUser.Email,
		CollectorRole:  service.CollectorroleSuperCollector,
		DisplayName:    devGothUser.NickName,
		AvatarUrl:      devCollector.AvatarUrl,
	}

	return nil
}
func (a *App) setupProductionModeCollectors() error {
	collectors, err := a.service.GetAllCollectors(context.Background())
	if err != nil {
		return err
	}
	a.collectors = make(map[string]service.Collector, len(collectors)+100)
	for _, v := range collectors {
		a.collectors[v.CollectorEmail] = v
	}

	// maps.Insert[string, service.Collector](a.collectors, func(yield func(string, service.Collector) bool) {

	// })
	return nil
}
