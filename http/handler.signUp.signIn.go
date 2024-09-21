package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"
)

func (a *App) signupPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := frontend.SignUpPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) signinPageHandler(w http.ResponseWriter, r *http.Request) {

	if err := frontend.SignInPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
