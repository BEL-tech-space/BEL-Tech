package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func (a *App) signInWithProviderAgent(w http.ResponseWriter, r *http.Request) {

	var provider string
	if provider = chi.URLParam(r, "provider"); provider == "" {
		a.clientError(w, http.StatusBadRequest, errors.New("provider not specified"))
		return
	}

	if provider != "google" && provider != "facebook" {
		a.clientError(w, http.StatusBadRequest, errors.New("invalid provider"))
		return
	}

	if a.config.Build != productionEnv { // not production
		http.Redirect(w, r, fmt.Sprintf("/agent/auth/%s/callback", provider), http.StatusTemporaryRedirect)
		return
	}

	q := r.URL.Query()
	q.Add("provider", provider)
	r.URL.RawQuery = q.Encode()

	// try to get the user without re-authenticating
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err == nil {
		if err = a.logInAgent(w, r, gothUser); err != nil {
			a.serverError(w, err)
			http.Redirect(w, r, "/agent/signin", http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/agent/dashboard", http.StatusTemporaryRedirect)
		return
	}

	gothic.BeginAuthHandler(w, r)
}

func (a *App) callbackHandlerAgent(w http.ResponseWriter, r *http.Request) {
	var (
		user = devGothUser
		err  error
	)

	if a.config.Build == productionEnv {
		if user, err = gothic.CompleteUserAuth(w, r); err != nil {
			a.serverError(w, err)
			http.Redirect(w, r, "/agent/signin", http.StatusTemporaryRedirect)
			return
		}
	}

	if err = a.logInAgent(w, r, user); err != nil {
		fmt.Println("err from loginUser", err.Error())
		a.serverError(w, err)
		http.Redirect(w, r, "/agent/signin", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/agent/dashboard", http.StatusTemporaryRedirect)
}

func (a *App) logoutAgent(w http.ResponseWriter, r *http.Request) {
	if err := gothic.Logout(w, r); err != nil {
		fmt.Println("error in goth.logout", err.Error())
	}

	session, err := a.agentSessionStore.Get(r, "beltech-agent")
	if err != nil {
		fmt.Println("error in get, logoutAgent", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	session.Options.MaxAge = -1
	session.Values = nil
	if err := session.Save(r, w); err != nil {
		fmt.Println("error in saving session, logOut", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (a *App) logInAgent(w http.ResponseWriter, r *http.Request, u goth.User) error {
	return a.createNewUserSessionAgent(w, r, u)
}

func (a *App) createNewUserSessionAgent(w http.ResponseWriter, r *http.Request, collector goth.User) error {
	session, err := a.agentSessionStore.New(r, "beltech-agent")
	if err != nil {
		fmt.Println("ERROR IN NEW", err.Error())
	}

	session.Values["email"] = collector.Email
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   7200, // Session expires in 2 hours
		HttpOnly: true, // Prevent JavaScript access
		Secure:   true, // Only send over HTTPS
		// SameSite: http.SameSiteStrictMode, // Only allow cookie for requests from the same site
	}

	// a.cacheStorage.SetDefault(user.Email, &user)

	return session.Save(r, w)
}
