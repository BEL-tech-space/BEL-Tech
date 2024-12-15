package http

import (
	accountPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)


func (a *App) dashboardPageHandler(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	if r.Header.Get("X-Content-Mode") == "partials" {
		if err := accountPages.DashboardMainPartial().Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accountPages.AccountDashboardPage(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

