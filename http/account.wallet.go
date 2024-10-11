package http

import (
	accountPages "beltech/frontend/pages/account"
	"beltech/service"
	"context"
	"net/http"
)

func (a *App) walletPageHandler(w http.ResponseWriter, r *http.Request) {
	var user *service.User
	userEmail := r.Context().Value(contextKeyEmail).(string)
	if x, found := a.cacheStorage.Get(userEmail); found {
		user = x.(*service.User)
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.WalletPage(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
