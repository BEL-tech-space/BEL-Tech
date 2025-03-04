package http

import (
	"context"
	"net/http"

	accPages "beltech/frontend/pages/account"

	"github.com/go-chi/chi/v5"
)

func (a *App) TipsPageHandler(w http.ResponseWriter, r *http.Request) {

	var wasteType string
	if wasteType = chi.URLParam(r, "wasteType"); wasteType == "" {
		http.Error(w, "WasteType not specified", http.StatusBadRequest)
		return
	}

	if wasteType != "glass" && wasteType != "metal" && wasteType != "paper" && wasteType != "plastics" {
		http.Error(w, "Invalid wasteType", http.StatusBadRequest)
		return
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.TipsOnRecyclingPage(wasteType, user).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

