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
		// http.Error(w, "Provider not specified", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	if provider != "google" && provider != "facebook" {
		// http.Error(w, "Invalid provider", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
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
		if err = a.logInUser(w, r, gothUser); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
			return
		}
	}

	if err = a.logInUser(w, r, user); err != nil {
		fmt.Println("err from loginUser", err.Error())
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/account/welcome", http.StatusTemporaryRedirect)
}

func (a *App) logoutAccount(w http.ResponseWriter, r *http.Request) {

	if err := gothic.Logout(w, r); err != nil {
		fmt.Println("error in goth.logout", err.Error())
	}

	session, err := a.accountSessionStore.Get(r, "beltech-account")
	if err != nil {
		fmt.Println("error in get, logoutAccount", err.Error())
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	session.Options.MaxAge = -1
	session.Values = nil
	if err := session.Save(r, w); err != nil {
		fmt.Println("error in saving session, logOut", err.Error())
		http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/account/login.html", http.StatusTemporaryRedirect)
}

func (a *App) logInUser(w http.ResponseWriter, r *http.Request, u goth.User) error {

	user, err := a.service.CreateOrUpdateUser(context.Background(), mapGothUserToServiceUser(u))
	if err != nil {
		fmt.Println("err came from create or update user:", err)
		return err
	}

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
		fmt.Println("ERROR IN NEW", err.Error())
	}

	session.Values["email"] = acc.Email
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86_400 * 14, // Session expires in 2 weeks
		HttpOnly: true,        // Prevent JavaScript access
		Secure:   true,        // Only send over HTTPS
		// SameSite: http.SameSiteLaxMode, // Only allow cookie for requests from the same site
	}

	a.accountCache.SetWithTTL(acc.Email, acc, 1, 86_400*14*time.Second)

	if a.config.Build != productionEnv {
		session.Options.MaxAge = 86_400
		a.accountCache.SetWithTTL(acc.Email, acc, 1, 86_400*time.Second)
	}

	return session.Save(r, w)
}

func mapGothUserToServiceUser(gothUser goth.User) service.Account {
	user := service.Account{}
	user.AccountID = gothUser.UserID
	user.Email = gothUser.Email
	user.FirstName = gothUser.FirstName
	user.LastName = gothUser.LastName
	user.DisplayName = gothUser.NickName
	user.AvatarUrl = gothUser.AvatarURL
	user.AccessToken = gothUser.AccessToken
	user.AccessTokenExpires = gothUser.ExpiresAt
	user.IDToken = gothUser.IDToken
	return user
}
