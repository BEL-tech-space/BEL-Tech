package http

import (
	"beltech/service"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func (a *App) signInWithProvider(w http.ResponseWriter, r *http.Request) {

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
		err = a.logInUser(w, r, gothUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("came to redirect")
		http.Redirect(w, r, "/beltech/account/dashboard", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("came to begin auth handler")
	gothic.BeginAuthHandler(w, r)
}

func (a *App) callbackHandler(w http.ResponseWriter, r *http.Request) {
	var (
		user goth.User
		err  error
	)

	if user, err = gothic.CompleteUserAuth(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = a.logInUser(w, r, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/beltech/account/dashboard", http.StatusTemporaryRedirect)
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}

func (a *App) logInUser(w http.ResponseWriter, r *http.Request, u goth.User) error {

	user := mapGothUserToServiceUser(u)
	if err := a.service.CreateOrUpdateUser(context.Background(), user); err != nil {
		return err
	}

	return a.createNewUserSession(w, r, user)
}

func (a *App) createNewUserSession(w http.ResponseWriter, r *http.Request, user service.User) error {

	session, err := a.accountSessionStore.Get(r, "beltech-session")
	if err != nil && !strings.Contains(err.Error(), "securecookie: the value is not valid") {
		fmt.Println("error in get session:", err)
		return err
	}

	session.Values["userID"] = user.UserID
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,                    // Session expires in 1 Hour
		HttpOnly: true,                    // Prevent JavaScript access
		Secure:   true,                    // Only send over HTTPS
		SameSite: http.SameSiteStrictMode, // Only allow cookie for requests from the same site
	}

	return session.Save(r, w)
}

func mapGothUserToServiceUser(gothUser goth.User) service.User {
	user := service.User{}
	user.UserID = gothUser.UserID
	user.Email = gothUser.Email
	user.FirstName = pgtype.Text{String: gothUser.FirstName, Valid: true}
	user.LastName = pgtype.Text{String: gothUser.LastName, Valid: true}
	user.NickName = pgtype.Text{String: gothUser.NickName, Valid: true}
	user.AvatarUrl = pgtype.Text{String: gothUser.AvatarURL, Valid: true}
	user.AccessToken = pgtype.Text{String: gothUser.AccessToken, Valid: true}
	user.AccessTokenExpires = pgtype.Timestamp{Time: gothUser.ExpiresAt, Valid: true}
	user.IDToken = pgtype.Text{String: gothUser.IDToken, Valid: true}
	return user
}
