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

		if a.config.Build != productionEnv {
			r.Get("/dashboard", a.dashboardPageHandler)
			r.Get("/location", a.locationPageHandler)
			r.Get("/wallet", a.walletPageHandler)
			r.Get("/spend", a.spendPageHandler)
			r.Get("/waste/{wasteType}", a.wasteHistoryHandler)
			r.Get("/tips/{wasteType}", a.TipsPageHandler)
			r.Get("/beneficiaries", a.beneficiariesPageHandler)
			r.Post("/airtime", a.airtimePageHandler)
			r.Get("/profile", a.ProfilePageHandler)

		}

		// this endpoint is reach from the signUp/login page
		// the aim is to get a fresh info of the sw from the server
		r.Get("/welcome", a.dashboardPageHandler)

		r.Get("/dashboard/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/dashboard/commonHeader-Partial.html", a.commonHeader)
		r.Get("/dashboard/mainOpeningTag-Partial.html", a.dashboardMainTagPartial)
		r.Get("/dashboard/swiperOne", a.dashboardSwiperOne)
		r.Get("/dashboard/echarts", a.dashboardDataViz)
		r.Get("/dashboard/swiperTwo", a.dashboardSwiperTwo)
		r.Get("/dashboard/footer-Partial.html", a.footerPartial)
		r.Get("/dashboard/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)
		r.Get("/dashboard/chart", a.dashboardEchart)

		r.Get("/scan", a.scanQRPageHandler)

		r.Get("/location/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/location/header-Partial.html", a.locationHeaderTagPartial)
		r.Get("/location/main", a.locationPageHandler)
		r.Get("/location/footer-Partial.html", a.footerPartial)
		r.Get("/location/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)

		r.Get("/wallet/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/wallet/commonHeader-Partial.html", a.commonHeader)
		r.Get("/wallet/mainOpeningTag-Partial.html", a.walletMainTagPartial)
		r.Get("/wallet/balance", a.walletBalance)
		r.Get("/wallet/options", a.walletOptionsPartial)
		r.Get("/wallet/coinData", a.walletcoinData)
		r.Get("/wallet/swiper", a.walletswiper)
		r.Get("/wallet/footer-Partial.html", a.footerPartial)
		r.Get("/wallet/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)
		r.Get("/wallet/livedata", a.walletLivecoinData)

		r.Get("/spend/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/spend/header-Partial.html", a.spendPageHeaderPartial)
		r.Get("/spend/mainOpeningTag-Partial.html", a.spendMainTagPartial)
		r.Get("/spend/balance", a.spendBalance)
		r.Get("/spend/otherhalf", a.spendOtherHalfPartial)
		r.Get("/spend/footer-Partial.html", a.footerPartial)
		r.Get("/spend/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)

		r.Get("/profile/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/profile/commonHeader-Partial.html", a.commonHeader)
		r.Get("/profile/mainOpeningTag-Partial.html", a.profileMainTagPartial)
		r.Get("/profile/displayIcon", a.profileDisplayIcon)
		r.Get("/profile/formPartial", a.profileFormPartial)
		r.Get("/profile/footer-Partial.html", a.footerPartial)
		r.Get("/profile/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)
		r.Post("/profile/changeDisplayName", a.changeProfileDisplayName)
		// r.Post("/profile/uploadAvatar", a.changeDisplayIcon)

		r.Get("/airtime/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/airtime/header-Partial.html", a.airtimePageHeaderPartial)
		r.Get("/airtime/mainOpeningTag-Partial.html", a.airtimeMainTagPartial)
		r.Get("/airtime/inputNumber", a.airtimeInputNumberPartial)
		r.Get("/airtime/beneficiaries", a.airtimeBeneficiaries)
		r.Get("/airtime/footer-Partial.html", a.footerPartial)
		r.Get("/airtime/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)

		r.Get("/beneficiaries/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/beneficiaries/header-Partial.html", a.beneficiariesPageHeaderPartial)
		r.Get("/beneficiaries/main", a.beneficiariesMain)
		r.Get("/beneficiaries/footer-Partial.html", a.footerPartial)
		r.Get("/beneficiaries/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)
		r.Post("/beneficiaries/addNew", a.addNewBeneficiary)
		r.Post("/beneficiaries/delete", a.deleteBeneficiary)

		r.Get("/waste/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/waste/commonHeader-Partial.html", a.commonHeader)
		r.Get("/waste/mainOpeningTag-Partial.html", a.wasteMainTagPartial)
		r.Get("/waste/quickInfo/{wasteType}", a.wasteQuickInfo)
		r.Get("/waste/filter", a.wasteFilterOptions)
		r.Get("/waste/history/{wasteType}", a.wasteHistory)
		r.Get("/waste/footer-Partial.html", a.footerPartial)
		r.Get("/waste/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)

		r.Get("/tips/openingHtmlTag-Partial.html", a.openingHThmlTagPartial)
		r.Get("/tips/tipsHeader-Partial.html", a.tipsHeaderPartial)
		r.Get("/tips/mainOpeningTag-Partial.html", a.tipsMainOpeningTag)
		r.Get("/tips/quickInfo/{wasteType}", a.wasteQuickInfo)
		r.Get("/tips/article/{wasteType}", a.tipsWasteArticle)
		r.Get("/tips/footer-Partial.html", a.footerPartial)
		r.Get("/tips/closingHtmlTag-Partial.html", a.closingHtmlTagPartial)

		r.Get("/history", a.historyPageHandler)
	})

	r.Get("/signup.html", a.signupPageHandler)
	r.Get("/login.html", a.signinPageHandler)

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

		// Add userID and accountEmail to the request context
		ctx := context.WithValue(r.Context(), contextKeyEmail, email.(string))
		r = r.WithContext(ctx)

		// Session is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}
