package http

import (
	"context"
	"net/http"
	"strings"

	accPages "beltech/frontend/pages/account"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func (a *App) openingHThmlTagPartial(w http.ResponseWriter, r *http.Request) {
	pageTitle := strings.Split(r.URL.Path, "/")[2]

	if err := accPages.HtmlAndBodyOpeningTagPartial(pageTitle).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) closingHtmlTagPartial(w http.ResponseWriter, r *http.Request) {
	pageTitle := strings.Split(r.URL.Path, "/")[2]

	if err := accPages.HtmlAndBodyClosingTagPartial(pageTitle).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) footerPartial(w http.ResponseWriter, r *http.Request) {
	pageTitle := strings.Split(r.URL.Path, "/")[2]

	if err := accPages.Footer(
		pageTitle,
		accPages.MainClosingTag(),
	).Render(context.TODO(), w); err != nil {
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
		// openingMainTagComp = accPages.WasteMainOpeningTag()
	}

	if err := openingMainTagComp.Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) wasteTipsHandler(w http.ResponseWriter, r *http.Request) {

	var wasteType string
	if wasteType = chi.URLParam(r, "wasteType"); wasteType == "" {
		http.Error(w, "WasteType not specified", http.StatusBadRequest)
		return
	}

	if wasteType != "glass" && wasteType != "metal" && wasteType != "paper" && wasteType != "plastics" {
		http.Error(w, "Invalid wasteType", http.StatusBadRequest)
		return
	}

	if err := accPages.TipsOnRecyclingPage(wasteType).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
