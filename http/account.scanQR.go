package http

import (
	accountPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) scanQRPageHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.ScanQRPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
