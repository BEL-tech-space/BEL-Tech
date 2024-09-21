package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) resetPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := frontend.ResetPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
