package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) signupPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", "/account/signup-Partial.html")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.SignUpPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) signinPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", "/account/signin-Partial.html")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.SignInPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
