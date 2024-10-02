package http

import (
	accountPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) splash_screenPageHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.SplashScreenPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
