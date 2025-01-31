package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) scanQRPageHandler(w http.ResponseWriter, r *http.Request) {
	accountEmail := r.Context().Value(contextKeyEmail).(string)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.ScanQRPage(accountEmail).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
