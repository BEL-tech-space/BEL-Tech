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
		r.Get("/logout/{provider}", a.logoutAccount)
	})

	r.Route("/", func(r chi.Router) {
		r.Use(a.accountAuthorization)

		// this endpoint is reach from the signUp/login page
		// the aim is to get a fresh info of the sw from the server
		r.Get("/welcome", a.dashboardPageHandler)

		r.Get("/dashboard", a.dashboardPageHandler)
		r.Get("/dashboard/chart", a.dashboardEchart)

		r.Get("/scan", a.scanQRPageHandler)
		r.Post("/confirmWaste", a.confirmWaste)
		r.Post("/submitWaste", a.submitWaste)

		r.Get("/history", a.historyPageHandler)

		r.Get("/location", a.locationPageHandler)

		r.Get("/wallet", a.walletPageHandler)
		r.Get("/wallet/livedata", a.walletLivecoinData)

		r.Get("/spend", a.spendPageHandler)

		r.Post("/profile/changeDisplayName", a.changeProfileDisplayName)
		// r.Post("/profile/uploadAvatar", a.changeDisplayIcon)

		r.Post("/airtime", a.airtimePageHandler)

		r.Get("/beneficiaries", a.beneficiariesPageHandler)
		r.Post("/beneficiaries/addNew", a.addNewBeneficiary)
		r.Post("/beneficiaries/delete", a.deleteBeneficiary)

		r.Get("/profile", a.ProfilePageHandler)

		r.Get("/notifications", a.getNotificationsHandler)
		r.Get("/nf", a.makeNotificationPageHandler)
		r.Post("/createNotification", a.createNotificationHandler)

		r.Get("/waste/{wasteType}", a.wasteHistoryHandler)
		r.Get("/tips/{wasteType}", a.TipsPageHandler)
	})

	r.Get("/signup.html", a.signupPageHandler)
	r.Get("/login.html", a.signinPageHandler)

	return r
}

func (a *App) accountAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the session
		session, err := a.accountSessionStore.Get(r, "beltech-account")
		if err != nil {
			fmt.Println("get err in auth", err.Error())
			http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
			return
		}

		// Check if the session lacks user information
		email, exist := session.Values["email"]
		if !exist {
			fmt.Println("SESSION EXPIRED")
			w.Header().Set("X-Session-Expired", "true")
			http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyEmail, email.(string))
		r = r.WithContext(ctx)

		// Session is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}
