package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) homePageHandler(w http.ResponseWriter, r *http.Request) {

	if err := frontend.HomePage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
