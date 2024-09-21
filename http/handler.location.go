package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) locationPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := frontend.LocationPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
