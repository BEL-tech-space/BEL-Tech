package http

import (
	"beltech/frontend"
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func (a *App) Routes(cfg EnvConfig) *chi.Mux {
	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Get("/", a.homePageHandler)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}", a.signInWithProvider)
		r.Get("/{provider}/callback", a.callbackHandler)
		r.Get("/logout/{provider}", a.logout)
	})

	r.Route("/beltech/account", func(r chi.Router) {
		r.Use(a.accountAuthorization)
		r.Get("/dashboard", a.accountDashboard)
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

	// r.Mount("/beltech/account", a.accountRouter(cfg))

	// r.Route("/beltech/agent", func(r chi.Router) {
	// 	r.Use(a.agentAuthorization)
	// 	r.Get("/qrcode/{agentID}", a.generateQRCode)
	// 	r.Post("/addAgent", a.addNewCollectorPost)
	// })

	// r.Get("/login", a.logInCollector)
	// r.Get("/logout", a.logOutCollector)
	// r.Get("/unauthorized", a.unauthorized)

	var staticFiles = fs.FS(frontend.AssetsDir)
	staticFs, _ := fs.Sub(staticFiles, "assets")
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(staticFs))))

	return r
}

func (a *App) accountRouter(cfg EnvConfig) http.Handler {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}", a.signInWithProvider)
		r.Get("/{provider}/callback", a.callbackHandler)
		r.Get("/logout/{provider}", a.logout)
	})

	r.Route("/", func(r chi.Router) {
		r.Use(a.accountAuthorization)
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

	return r
}

const (
	beltechAccount string = "beltech-account"
	beltechAgent string = "beltech-agent"
)


func (a *App) accountAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the session
		session, err := a.accountSessionStore.Get(r, beltechAccount)
		if err != nil {
			if err.Error() == "securecookie: the value is not valid" {
				log.Println("err in get session auth", err)
				// a.setFlashMessage(w, r, "invalid session, let get a new one...")
				http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
				return
			}
			
			fmt.Println("new error", err.Error())	
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
			return
		}

		// Check if the session lacks user information
		userID, exist := session.Values["userID"]
		if !exist {
			log.Println("session did not exit")
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
			return
		}

		var userEmail string
		if email, exist := session.Values["email"]; exist {
			userEmail = email.(string)
		}

		// Add userID and userEmail to the request context
		ctx := context.WithValue(r.Context(), contextKeyEmail, userEmail)
		ctx = context.WithValue(ctx, contextKeyAgentID, userID.(string))

		r = r.WithContext(ctx)

		// Session is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}

type contextKey string

const (
	contextKeyEmail      contextKey = "email"
	contextUserAvatarURL contextKey = "avatarURL"
	contextKeyAgentID    contextKey = "agentID"
	contextKeyUserID     contextKey = "userID"
	contextKeyRole       contextKey = "role"
)

func (a *App) agentAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the session
		session, err := a.accountSessionStore.Get(r, "beltech-agent")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if the session lacks user information
		email, exist := session.Values["email"]
		if !exist {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		agentEmail := email.(string)

		var (
			agentID        uuid.UUID
			role           string
			foundCollector bool
		)

		for _, collector := range a.collectors {
			if collector.Email != agentEmail {
				continue
			}

			agentID = collector.AgentID
			role = collector.Role
			foundCollector = true
			break
		}

		if !foundCollector {
			http.Redirect(w, r, "/unauthorized", http.StatusUnauthorized)
			return
		}

		// Add agentID and agentID to the request context
		ctx := context.WithValue(r.Context(), contextKeyEmail, agentEmail)
		ctx = context.WithValue(ctx, contextKeyAgentID, agentID)
		ctx = context.WithValue(ctx, contextKeyRole, role)

		r = r.WithContext(ctx)

		// Session is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
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
