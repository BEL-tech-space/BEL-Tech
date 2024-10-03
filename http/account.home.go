package http

import (
	accountPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) accountDashboard(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.AccountDashboardPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
