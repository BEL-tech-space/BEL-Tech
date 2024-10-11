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

	fmt.Println("at BeginAuthHandler")
	gothic.BeginAuthHandler(w, r)
	fmt.Println("Past BeginAuthHandler")
}

func (a *App) callbackHandler(w http.ResponseWriter, r *http.Request) {
	var (
		user goth.User
		err  error
	)

	fmt.Println("in callback")
	
	if user, err = gothic.CompleteUserAuth(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("past ")

	if err = a.logInUser(w, r, user); err != nil {
		fmt.Println("err from loginUser", err.Error())
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

	user, err := a.service.CreateOrUpdateUser(context.Background(), mapGothUserToServiceUser(u)); 
	if err != nil {
		fmt.Println("err came from create or update user:", err)
		return err
	}

	return a.createNewUserSession(w, r, user)
}

func (a *App) createNewUserSession(w http.ResponseWriter, r *http.Request, user service.User) error {

	session, err := a.accountSessionStore.New(r, "beltech-account")
	if err != nil {
		fmt.Println("ERROR IN NEW", err.Error())
	}

	session.Values["userID"] = user.UserID
	session.Values["email"] = user.Email
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,                    // Session expires in 1 Hour
		HttpOnly: true,                    // Prevent JavaScript access
		Secure:   true,                    // Only send over HTTPS
		SameSite: http.SameSiteStrictMode, // Only allow cookie for requests from the same site
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
	user.AvatarUrl = gothUser.AvatarURL
	user.AccessToken = gothUser.AccessToken
	user.AccessTokenExpires = gothUser.ExpiresAt
	user.IDToken = gothUser.IDToken
	return user
}
