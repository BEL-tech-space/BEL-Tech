package http

import (
	"context"
	"errors"
	"net/http"

	accPages "beltech/frontend/pages/account"

	"github.com/go-chi/chi/v5"
)

func (a *App) TipsPageHandler(w http.ResponseWriter, r *http.Request) {

	var wasteType string
	if wasteType = chi.URLParam(r, "wasteType"); wasteType == "" {
		a.clientError(w, http.StatusBadRequest, errors.New("wasteType not specified"))
		return
	}

	if wasteType != "glass" && wasteType != "metal" && wasteType != "paper" && wasteType != "plastics" {
		a.clientError(w, http.StatusBadRequest, errors.New("invalid wasteType"))
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
		a.serverError(w, err)
		return
	}
}
