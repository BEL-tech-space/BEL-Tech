package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) password_resetPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := frontend.PasswordResetPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
