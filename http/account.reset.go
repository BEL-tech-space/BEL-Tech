package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) resetPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := accPages.ResetPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
