package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) agentRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}", a.signInWithProviderAgent)
		r.Get("/{provider}/callback", a.callbackHandler)
		// r.Get("/logout/{provider}", a.logout)
	})

	r.Route("/", func(r chi.Router) {
		r.Use(a.agentAuthorization)
		r.Get("/dashboard", a.agentDashboard)
		r.Get("/collect", a.collectPage)
		r.Get("/history", a.historyPage)
		r.Get("/profile", a.profilePage)
		r.Post("/profile/updateDisplayName", a.upUpdateCollectorDisplayName)
		r.Get("/signout.html", a.logOutCollector)

		r.Post("/qrcode/{agentID}", a.generateQRCode)
		r.Get("/addNewCollector", a.addNewCollector)
		r.Post("/addNewCollector", a.addNewCollectorPost)
		r.Get("/waste/{wasteType}", a.agentWasteHistoryHandler)

	})

	r.Get("/logout", a.logoutAgent)
	r.Get("/login.html", a.logInCollector)
	r.Get("/unauthorized", a.unauthorized)

	return r
}

func (a *App) agentAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the session
		session, err := a.agentSessionStore.Get(r, "beltech-agent")
		if err != nil {
			http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
			return
		}

		// Check if the session lacks user information
		email, exist := session.Values["email"]
		if !exist {
			http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
			return
		}

		agentEmail := email.(string)

		collector, ok := a.collectors[agentEmail]
		if !ok {
			http.Redirect(w, r, "/agent/unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyEmail, agentEmail)
		ctx = context.WithValue(ctx, contextKeyAgentID, collector.AgentID)
		ctx = context.WithValue(ctx, contextKeyRole, collector.CollectorRole)

		r = r.WithContext(ctx)

		// Session is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}
