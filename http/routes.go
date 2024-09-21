package http

import (
	"beltech/frontend"
	"beltech/utils"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) Routes(cfg utils.EnvConfig) *chi.Mux {
	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}", a.signInWithProvider)
		r.Get("/{provider}/callback", a.callbackHandler)
		r.Get("/logout/{provider}", a.logout)
	})

	r.Route("/beltech", func(r chi.Router) {
		r.Use(a.authorize)
		r.Get("/dashboard", a.homePageHandler)
		r.Get("/scan", a.homePageHandler)
		r.Get("/bins", a.homePageHandler)
		r.Get("/wallet", a.homePageHandler)
		r.Get("/exchange", a.homePageHandler)
	})

	r.Get("/signup", a.signupPageHandler)
	r.Get("/signin", a.signinPageHandler)
	r.Get("/verify", a.verifyPageHandler)
	r.Get("/reset", a.resetPageHandler)
	r.Get("/verify_success", a.verify_successPageHandler)
	r.Get("/password_reset", a.password_resetPageHandler)
	r.Get("/home", a.splash_screenPageHandler)

	var staticFiles = fs.FS(frontend.AssetsDir)
	staticFs, _ := fs.Sub(staticFiles, "assets")
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(staticFs))))

	return r
}
