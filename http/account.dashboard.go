package http

import (
	accountPages "beltech/frontend/pages/account"
	"context"
	"fmt"
	"net/http"
)

func (a *App) accountDashboard(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		fmt.Println("error occured here")
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.AccountDashboardPage(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) accountProfilePage(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.AccountProfile(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
