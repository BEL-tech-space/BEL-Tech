package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) historyPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("X-Content-Mode") == "partials" {
		if err := accPages.HistoryMain().Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accPages.HistoryPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) historyPageHeaderPartial(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.HistoryHeaderPartialComp(
		accPages.BeforeHeaderTagPartial(
			"History"),
		accPages.HistoryMainOpeningTagComp()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) historyList(w http.ResponseWriter, r *http.Request) {

	if err := accPages.HistoryMain().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
