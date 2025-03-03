package http

import (
	"beltech/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/pkg/errors"
)

func (a *App) signInWithProviderAgent(w http.ResponseWriter, r *http.Request) {

	var provider string
	if provider = chi.URLParam(r, "provider"); provider == "" {
		http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
		return
	}

	if provider != "google" && provider != "facebook" {
		http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
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
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		if err = a.logInAgent(w, r, gothUser); err != nil {
			http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
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
			http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
			return
		}
	}

	if err = a.logInAgent(w, r, user); err != nil {
		http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/agent/dashboard", http.StatusTemporaryRedirect)
}

func (a *App) logoutAgent(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)

	session, err := a.agentSessionStore.Get(r, "beltech-agent")
	if err != nil || session == nil {
		http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
		return
	}

	session.Options.MaxAge = -1
	session.Values = nil
	if err := session.Save(r, w); err != nil {
		http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/agent/login.html", http.StatusTemporaryRedirect)
}

func (a *App) logInAgent(w http.ResponseWriter, r *http.Request, u goth.User) error {

	collector, ok := a.collectors[u.Email]
	if !ok {
		return errors.New("not a collector")
	}

	return a.createNewUserSessionAgent(w, r, collector)
}

func (a *App) createNewUserSessionAgent(w http.ResponseWriter, r *http.Request, collector service.AgentCollector) error {
	session, err := a.agentSessionStore.New(r, "beltech-agent")
	if err != nil {
		fmt.Println("ERROR IN NEW", err.Error())
	}

	session.Values["email"] = collector.CollectorEmail
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   7200, // Session expires in 2 hours
		HttpOnly: true, // Prevent JavaScript access
		Secure:   true, // Only send over HTTPS
		// SameSite: http.SameSiteStrictMode, // Only allow cookie for requests from the same site
	}

	// TODO: cache collector info

	return session.Save(r, w)
}
