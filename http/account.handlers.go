package http

import (
	"context"
	"net/http"
	"strings"

	accPages "beltech/frontend/pages/account"

	"github.com/a-h/templ"
)

func (a *App) openingHThmlTagPartial(w http.ResponseWriter, r *http.Request) {
	pageTitle := strings.Split(r.URL.Path, "/")[2]

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.HtmlAndBodyOpeningTagPartial(pageTitle).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) closingHtmlTagPartial(w http.ResponseWriter, r *http.Request) {
	pageTitle := strings.Split(r.URL.Path, "/")[2]

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.HtmlAndBodyClosingTagPartial(pageTitle).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) footerPartial(w http.ResponseWriter, r *http.Request) {
	pageTitle := strings.Split(r.URL.Path, "/")[2]

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.Footer(pageTitle, accPages.MainClosingTag()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) commonHeader(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.CommonHeader(
		user).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) mainOpeningTagPartial(w http.ResponseWriter, r *http.Request) {
	pageTitle := strings.Split(r.URL.Path, "/")[2]

	var openingMainTagComp templ.Component
	switch pageTitle {
	case "dashboard":
		openingMainTagComp = accPages.DashboardMainOpeningTag()
	case "wallet":
		openingMainTagComp = accPages.WalletMainOpeningTag()
	case "spend":
		openingMainTagComp = accPages.SpendMainOpeningTag()
	case "profile":
		openingMainTagComp = accPages.ProfileMainOpeningTag()
	case "waste":
		openingMainTagComp = accPages.WasteMainOpeninTag()
	}

	if err := openingMainTagComp.Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
