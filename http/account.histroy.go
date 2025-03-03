package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"fmt"
	"net/http"
)

func (a *App) historyPageHandler(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	wasteLog, err := a.service.GetWasteLogByAccountPaginated(
		r.Context(),
		accountEmail,
		30,
		0,
	)
	if err != nil {
		a.serverError(w, fmt.Errorf("error getting wastelog: %w", err))
	}

	if r.Header.Get("X-Content-Mode") == "partials" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := accPages.HistoryMain(wasteLog).Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.HistoryPage(wasteLog).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) historyPageHeaderPartial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.HistoryHeaderPartial(
		accPages.BeforeHeaderTagPartial(
			"History"),
		accPages.HistoryMainOpeningTag()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) historyList(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	wasteLog, err := a.service.GetWasteLogByAccountPaginated(
		r.Context(),
		accountEmail,
		30,
		0,
	)
	if err != nil {
		a.serverError(w, fmt.Errorf("error getting wastelog: %w", err))
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.HistoryMain(wasteLog).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
