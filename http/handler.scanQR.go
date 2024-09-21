package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) scanQRPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := frontend.ScanQRPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
