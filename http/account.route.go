package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) accountRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}", a.signInWithProviderAccount)
		r.Get("/{provider}/callback", a.callbackHandlerAccount)
		r.Get("/logout/{provider}", a.logout)
	})

	r.Route("/", func(r chi.Router) {
		r.Use(a.accountAuthorization)

		r.Get("/dashboard", a.accountDashboard)
		r.Get("/scan", a.scanQRPageHandler)
		r.Get("/location", a.locationPageHandler)
		r.Get("/wallet", a.walletPageHandler)
		r.Get("/spend", a.exchangePageHandler)
		r.Get("/profile", a.accountProfilePage)
		r.Post("/airtime", a.rechargeAirtime)
		r.Get("/beneficiaries", a.airtimeBeneficiaries)
		r.Get("/waste/{wasteType}", a.wasteHistoryHandler)
		r.Get("/tips/{wasteType}", a.wasteTipsHandler)
		r.Get("/history", a.allWasteHistoryHandler)

	})

	r.Get("/signup", a.signupPageHandler)
	r.Get("/signin", a.signinPageHandler)
	r.Get("/verify", a.verifyPageHandler)
	r.Get("/reset", a.resetPageHandler)
	r.Get("/verify_success", a.verify_successPageHandler)

	return r
}

func (a *App) accountAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the session
		session, err := a.accountSessionStore.Get(r, "beltech-account")
		if err != nil {
			fmt.Println("get err in auth", err.Error())
			http.Redirect(w, r, "/account/signin", http.StatusTemporaryRedirect)
			return
		}

		// Check if the session lacks user information
		email, exist := session.Values["email"]
		if !exist {
			fmt.Println("email not in session, weird")
			http.Redirect(w, r, "/account/signin", http.StatusTemporaryRedirect)
			return
		}

		// Add userID and userEmail to the request context
		ctx := context.WithValue(r.Context(), contextKeyEmail, email.(string))
		r = r.WithContext(ctx)

		// Session is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}
