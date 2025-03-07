package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) airtimePageHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, err)
		return
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	result, err := a.service.GetBeneficiaries(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.AirtimePage(user, result).Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
