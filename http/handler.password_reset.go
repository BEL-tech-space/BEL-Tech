package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) password_resetPageHandler(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := frontend.PasswordResetPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
