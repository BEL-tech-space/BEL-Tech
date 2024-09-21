package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) splash_screenPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := frontend.SplashScreenPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
