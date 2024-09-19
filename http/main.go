package http

import (
	"beltech/service"
	"beltech/utils"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type App struct {
	service      *service.Service
	sessionStore sessions.Store
}

func NewApp(cfg utils.EnvConfig) *App {

	store := sessions.NewCookieStore([]byte(cfg.CookieSecret))
	store.MaxAge(86400 * 30) // 30 days
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = true
	gothic.Store = store

	goth.UseProviders(
		google.New(cfg.LoginProvider.Google.ClientKey, cfg.LoginProvider.Google.Secret, "http://localhost:8090/auth/google/callback"),
		// facebook.New(cfg.LoginProvider.Facebook.ClientKey, cfg.LoginProvider.Facebook.Secret, "http://localhost:8090/auth/facebook/callback"),
	)

	app := App{}
	app.service = service.NewService()

	return &app
}
