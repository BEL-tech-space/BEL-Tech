package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) locationPageHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := accPages.LocationPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
