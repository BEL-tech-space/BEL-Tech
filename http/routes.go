package http

import (
	"beltech/frontend"
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
	r.Mount("/agent", a.agentRouter())

	var staticFiles = fs.FS(frontend.AssetsDir)
	staticFs, _ := fs.Sub(staticFiles, "assets")
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(staticFs))))

	return r
}

// func (a *App) setFlashMessage(w http.ResponseWriter, r *http.Request, message string) {
//     session, _ := a.accountSessionStore.Get(r, "beltech-account")
//     session.AddFlash(message)
//     session.Save(r, w)
// }

// func (a *App) getFlashMessage(w http.ResponseWriter, r *http.Request) string {
//     session, _ := a.accountSessionStore.Get(r, "beltech-account")
//     if flashes := session.Flashes(); len(flashes) > 0 {
//         session.Save(r, w)
//         return flashes[0].(string)
//     }
//     return ""
// }
