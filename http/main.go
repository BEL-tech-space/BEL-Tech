package http

import (
	"beltech/service"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type App struct {
	service *service.Service
}

func NewApp() *App {

	key := "" //your client secret
	maxAge := 86400 * 30 // 30 days
	isProd := true       // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
    	google.New("",
		 key, 
		 "http://localhost:8090/auth/google/callback"),
)
	return &App{}
}