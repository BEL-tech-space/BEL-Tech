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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.HistoryPage(wasteLog).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



