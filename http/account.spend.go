package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) spendPageHandler(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.SpendPage(user).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
