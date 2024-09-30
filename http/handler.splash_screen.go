package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) splash_screenPageHandler(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := frontend.SplashScreenPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
