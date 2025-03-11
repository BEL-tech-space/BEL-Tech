package http

import (
	"beltech/service"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/pkg/errors"
)

func (a *App) signInWithProviderAgent(w http.ResponseWriter, r *http.Request) {
	a.signInWithProvider(w, r, "agent")
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
		fmt.Printf("agent %s not found in %+v \n", u.Email, a.collectors)
		return errors.New("not a collector")
	}

	return a.createNewUserSessionAgent(w, r, collector)
}

func (a *App) createNewUserSessionAgent(w http.ResponseWriter, r *http.Request, collector service.Collector) error {
	session, err := a.agentSessionStore.New(r, "beltech-agent")
	if err != nil {
		fmt.Println("ERROR IN NEW", err.Error())
	}

	session.Values["email"] = collector.CollectorEmail
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86_400 * 2, // Session expires in 2 weeks
		HttpOnly: true,       // Prevent JavaScript access
		Secure:   true,       // Only send over HTTPS
		// SameSite: http.SameSiteStrictMode, // Only allow cookie for requests from the same site
	}

	// TODO: cache collector info

	return session.Save(r, w)
}
