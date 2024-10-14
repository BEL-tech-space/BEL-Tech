package http

import (
	"beltech/service"
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

var (
	devGothUser = goth.User{
		UserID:    "1234567890",
		Email:     "giwaoluatobi@gmail.com",
		AvatarURL: "https://lh3.googleusercontent.com/a-/ALV-UjWhunlRgNtwAAdFNTue7WN3k01azp8GY0qJHQqsdfXBB26X40qIIw=s96-c",
		NickName:  "Tobias",
	}

	devServiceUser = service.User{
		UserID:    "1234567890",
		Email:     "giwaoluatobi@gmail.com",
		AvatarURL: "https://lh3.googleusercontent.com/a-/ALV-UjWhunlRgNtwAAdFNTue7WN3k01azp8GY0qJHQqsdfXBB26X40qIIw=s96-c",
		NickName:  "Tobias",
	}
)

func (a *App) signInWithProviderAccount(w http.ResponseWriter, r *http.Request) {

	var provider string
	if provider = chi.URLParam(r, "provider"); provider == "" {
		http.Error(w, "Provider not specified", http.StatusBadRequest)
		return
	}

	if provider != "google" && provider != "facebook" {
		http.Error(w, "Invalid provider", http.StatusBadRequest)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err = a.logInUser(w, r, user); err != nil {
		fmt.Println("err from loginUser", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/account/dashboard", http.StatusTemporaryRedirect)
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (a *App) logInUser(w http.ResponseWriter, r *http.Request, u goth.User) error {
	var (
		user = devServiceUser
		err  error
	)

	if a.config.Build == productionEnv {
		if user, err = a.service.CreateOrUpdateUser(context.Background(), mapGothUserToServiceUser(u)); err != nil {
			fmt.Println("err came from create or update user:", err)
			return err
		}
	}

	return a.createNewUserSession(w, r, user)
}

func (a *App) createNewUserSession(w http.ResponseWriter, r *http.Request, user service.User) error {
	session, err := a.accountSessionStore.New(r, "beltech-account")
	if err != nil {
		fmt.Println("ERROR IN NEW", err.Error())
	}

	session.Values["email"] = user.Email
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600, // Session expires in 1 Hour
		HttpOnly: true, // Prevent JavaScript access
		Secure:   true, // Only send over HTTPS
		// SameSite: http.SameSiteLaxMode, // Only allow cookie for requests from the same site
	}

	a.cacheStorage.SetDefault(user.Email, &user)

	return session.Save(r, w)
}

func mapGothUserToServiceUser(gothUser goth.User) service.User {
	user := service.User{}
	user.UserID = gothUser.UserID
	user.Email = gothUser.Email
	user.FirstName = gothUser.FirstName
	user.LastName = gothUser.LastName
	user.NickName = gothUser.NickName
	user.AvatarURL = gothUser.AvatarURL
	user.AccessToken = gothUser.AccessToken
	user.AccessTokenExpires = gothUser.ExpiresAt
	user.IDToken = gothUser.IDToken
	return user
}
