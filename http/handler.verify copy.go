package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) verifyPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := frontend.VerifyPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
