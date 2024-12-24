package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) scanQRPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := accPages.ScanQRPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
