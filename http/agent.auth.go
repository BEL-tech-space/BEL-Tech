package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func (a *App) signInWithProviderAgent(w http.ResponseWriter, r *http.Request) {

	var provider string
	if provider = chi.URLParam(r, "provider"); provider == "" {
		http.Error(w, "Provider not specified", http.StatusBadRequest)
		return
	}

	if provider != "google" && provider != "facebook" {
		http.Error(w, "Invalid provider", http.StatusBadRequest)
		return
	}

	q := r.URL.Query()
	q.Add("provider", provider)
	r.URL.RawQuery = q.Encode()

	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		err = a.logInUserAgent(w, r, gothUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("came to redirect")
		http.Redirect(w, r, "/beltech/agent/dashboard", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("came to begin auth handler")
	gothic.BeginAuthHandler(w, r)
}

func (a *App) callbackHandlerAgent(w http.ResponseWriter, r *http.Request) {
	var (
		user goth.User
		err  error
	)

	if user, err = gothic.CompleteUserAuth(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = a.logInUserAgent(w, r, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/beltech/agent/dashboard", http.StatusTemporaryRedirect)
}

func (a *App) logoutAgent(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}

func (a *App) logInUserAgent(w http.ResponseWriter, r *http.Request, u goth.User) error {
	// user := mapGothUserToServiceUser(u)
	return a.createNewUserSessionAgent(w, r, u)
}

func (a *App) createNewUserSessionAgent(w http.ResponseWriter, r *http.Request, collector goth.User) error {

	session, err := a.accountSessionStore.Get(r, "beltech-session")
	if err != nil && !strings.Contains(err.Error(), "securecookie: the value is not valid") {
		fmt.Println("error in get session:", err)
		return err
	}

	session.Values["email"] = collector.Email
	// session.Values["agentID"] = collector.AgentID

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   7200,                    // Session expires in 2 hours
		HttpOnly: true,                    // Prevent JavaScript access
		Secure:   true,                    // Only send over HTTPS
		SameSite: http.SameSiteStrictMode, // Only allow cookie for requests from the same site
	}

	return session.Save(r, w)
}

