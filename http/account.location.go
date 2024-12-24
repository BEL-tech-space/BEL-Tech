package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
	"strings"
)

func (a *App) locationPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("X-Content-Mode") == "partials" {
		if err := accPages.LocationMain().Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accPages.LocationPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) locationHeaderTagPartial(w http.ResponseWriter, r *http.Request) {

	if err := accPages.LocationHeaderPartial(accPages.BeforeHeaderTagPartial(strings.Split(r.URL.Path, "/")[2])).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
