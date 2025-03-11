package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (a *App) signInWithProvider(w http.ResponseWriter, r *http.Request, userType string) {
	var provider string
	if provider = chi.URLParam(r, "provider"); provider == "" {
		http.Redirect(w, r, fmt.Sprintf("/%s/login.html", userType), http.StatusTemporaryRedirect)
		return
	}

	if provider != "google" && provider != "facebook" {
		http.Redirect(w, r, fmt.Sprintf("/%s/login.html", userType), http.StatusTemporaryRedirect)
		return
	}

	q := r.URL.Query()
	q.Add("provider", provider)
	q.Add("state", userType)
	r.URL.RawQuery = q.Encode()

	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		if userType == "agent" {
			if err = a.logInAgent(w, r, gothUser); err != nil {
				http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
				return
			}

			http.Redirect(w, r, "/agent/dashboard", http.StatusTemporaryRedirect)
			return
		}

		if err = a.logInAccount(w, r, gothUser); err != nil {
			http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
			return
		}
	}

	gothic.BeginAuthHandler(w, r)
}

func (a *App) callbackHandler(w http.ResponseWriter, r *http.Request) {

	userType := r.URL.Query().Get("state")
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/%s/login.html", userType), http.StatusTemporaryRedirect)
		return
	}

	if userType == "agent" {
		if err = a.logInAgent(w, r, user); err != nil {
			http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/agent/dashboard", http.StatusTemporaryRedirect)
		return
	}

	if err = a.logInAccount(w, r, user); err != nil {
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/account/welcome", http.StatusTemporaryRedirect)
}
