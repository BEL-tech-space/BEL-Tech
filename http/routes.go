package http

import (
	frontend "beltech/frontend"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) Routes() *chi.Mux {
	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Get("/", a.homePageHandler)

	r.Mount("/account", a.accountRouter())
	if a.config.Build != productionEnv {
		r.Mount("/agent", a.agentRouter())
	}

	var staticFiles = fs.FS(frontend.AssetsDir)
	staticFs, _ := fs.Sub(staticFiles, "public")

	r.Handle("/public/sw.js", http.HandlerFunc(a.sw))
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.FS(staticFs))))

	return r
}
