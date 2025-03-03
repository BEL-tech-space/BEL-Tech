package http

import (
	"beltech/service"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

var (
	devGothUser = goth.User{
		UserID:    "1234567890",
		Email:     "giwaoluwatobi@gmail.com",
		AvatarURL: "https://lh3.googleusercontent.com/a-/ALV-UjWhunlRgNtwAAdFNTue7WN3k01azp8GY0qJHQqsdfXBB26X40qIIw=s96-c",
		NickName:  "Tobias",
	}
)

func (a *App) signInWithProviderAccount(w http.ResponseWriter, r *http.Request) {

	var provider string
	if provider = chi.URLParam(r, "provider"); provider == "" {
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	if provider != "google" && provider != "facebook" {
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	if a.config.Build != productionEnv {
		http.Redirect(w, r, fmt.Sprintf("/account/auth/%s/callback", provider), http.StatusTemporaryRedirect)
		return
	}

	q := r.URL.Query()
	q.Add("provider", provider)
	r.URL.RawQuery = q.Encode()

	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		if err = a.logInAccount(w, r, gothUser); err != nil {
			http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/account/dashboard", http.StatusTemporaryRedirect)
		return
	}

	gothic.BeginAuthHandler(w, r)
}

func (a *App) callbackHandlerAccount(w http.ResponseWriter, r *http.Request) {
	var (
		user = devGothUser
		err  error
	)

	if a.config.Build == productionEnv {
		if user, err = gothic.CompleteUserAuth(w, r); err != nil {
			http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
			return
		}
	}

	if err = a.logInAccount(w, r, user); err != nil {
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/account/welcome", http.StatusTemporaryRedirect)
}

func (a *App) logInAccount(w http.ResponseWriter, r *http.Request, u goth.User) error {

	user, err := a.service.CreateOrUpdateUser(context.Background(), mapGothUserToServiceUser(u))
	if err != nil {
		fmt.Println("err came from create or update user:", err)
		return err
	}

	fmt.Printf("\n\non login--- %+v\n\n", user)

	fmt.Println(user.IsNew, "is new")

	if user.IsNew {
		http.SetCookie(w, &http.Cookie{
			Name:  "isNewAccount",
			Value: "true",
			Path:  "/",
		})
	}

	return a.createNewUserSession(w, r, user)
}

func (a *App) createNewUserSession(w http.ResponseWriter, r *http.Request, acc service.Account) error {
	session, err := a.accountSessionStore.New(r, "beltech-account")
	if err != nil {
		fmt.Printf("error creating new session: error: %s\n", err.Error())
	}

	session.Values["email"] = acc.AccountEmail
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86_400 * 14, // Session expires in 2 weeks
		HttpOnly: true,        // Prevent JavaScript access
		Secure:   true,        // Only send over HTTPS
		// SameSite: http.SameSiteLaxMode, // Only allow cookie for requests from the same site
	}

	a.accountCache.SetWithTTL(acc.AccountEmail, acc, 1, 86_400*14*time.Second)

	if a.config.Build != productionEnv { // dev mode
		session.Options.MaxAge = 86_400
		a.accountCache.SetWithTTL(acc.AccountEmail, acc, 1, 86_400*time.Second)
	}

	if err := session.Save(r, w); err != nil {
		return fmt.Errorf("error saving new session: error: %w", err)
	}

	return nil
}

func mapGothUserToServiceUser(gothUser goth.User) service.Account {
	user := service.Account{}
	user.AccountID = gothUser.UserID
	user.AccountEmail = gothUser.Email
	user.FirstName = gothUser.FirstName
	user.LastName = gothUser.LastName
	user.DisplayName = gothUser.NickName
	user.AvatarUrl = gothUser.AvatarURL
	user.AccessToken = gothUser.AccessToken
	user.AccessTokenExpires = gothUser.ExpiresAt
	user.IDToken = gothUser.IDToken
	return user
}
func (a *App) logoutAccount(w http.ResponseWriter, r *http.Request) {

	gothic.Logout(w, r)

	session, err := a.accountSessionStore.Get(r, "beltech-account")
	if err != nil || session == nil {
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	session.Options.MaxAge = -1
	session.Values = nil
	if err := session.Save(r, w); err != nil {
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
}
