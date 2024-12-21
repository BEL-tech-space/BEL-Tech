package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
	"strings"
)

func (a *App) locationPageHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.LocationHeaderPartialComp(accPages.BeforeHeaderTagPartial(strings.Split(r.URL.Path, "/")[2])).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
