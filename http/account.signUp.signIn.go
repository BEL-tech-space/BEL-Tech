package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) signupPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := accPages.SignUpPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) signinPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := accPages.SignInPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
