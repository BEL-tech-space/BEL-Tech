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
	r.Use(middleware.Compress(5))

	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}", a.signInWithProvider)
		r.Get("/{provider}/callback", a.callbackHandler)
		r.Get("/logout/{provider}", a.logout)
	})

	r.Route("/beltech", func(r chi.Router) {
		r.Use(a.authorize)
		r.Get("/dashboard", a.homePageHandler)
		r.Get("/scan", a.scanQRPageHandler)
		r.Get("/bins", a.locationPageHandler)
		r.Get("/wallet", a.walletPageHandler)
		r.Get("/exchange", a.exchangePageHandler)
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

func (a *App) authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the session
		session, err := a.sessionStore.Get(r, "beltech-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if the session is new or if it lacks user information
		if session.IsNew || session.Values["userID"] == nil {
			http.Redirect(w, r, "/signup", http.StatusUnauthorized)
			return
		}

		// Session is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}