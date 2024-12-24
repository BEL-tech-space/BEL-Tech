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

		r.Get("/dashboard/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/dashboard/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)
		r.Get("/dashboard/header-Partial.html", a.dashboardHeaderPartial)
		r.Get("/dashboard/footer-Partial.html", a.footerPartial)
		r.Get("/dashboard/chart", a.dashboardEchart)
		r.Get("/dashboard", a.dashboardPageHandler)

		r.Get("/scan", a.scanQRPageHandler)

		r.Get("/location/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/location/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)
		r.Get("/location/header-Partial.html", a.locationHeaderTagPartial)
		r.Get("/location/footer-Partial.html", a.footerPartial)
		r.Get("/location", a.locationPageHandler)

		r.Get("/wallet/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/wallet/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)
		r.Get("/wallet/header-Partial.html", a.walletHeaderPartial)
		r.Get("/wallet/footer-Partial.html", a.footerPartial)
		r.Get("/wallet", a.walletPageHandler)
		r.Get("/coinData", a.coinDataHandler)

		r.Get("/spend/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/spend/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)
		r.Get("/spend/header-Partial.html", a.spendPageHeaderPartial)
		r.Get("/spend/footer-Partial.html", a.footerPartial)
		r.Get("/spend", a.spendPageHandler)

		r.Get("/profile", a.ProfilePageHandler)
		r.Post("/profile/displayName", a.changeProfileDisplayName)

		r.Post("/airtime", a.airtimePageHandler)

		r.Get("/history", a.historyPageHandler)

		r.Get("/beneficiaries", a.beneficiariesPageHandler)
		r.Post("/beneficiaries/addNew", a.addNewBeneficiary)

		r.Get("/waste/{wasteType}", a.wasteHistoryHandler)
		r.Get("/tips/{wasteType}", a.wasteTipsHandler)

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

		// Add userID and accountEmail to the request context
		ctx := context.WithValue(r.Context(), contextKeyEmail, email.(string))
		r = r.WithContext(ctx)

		// Session is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}
