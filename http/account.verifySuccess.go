package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) verify_successPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := accPages.VerifySuccessPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
