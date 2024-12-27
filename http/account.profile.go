package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"fmt"
	"net/http"
)

func (a *App) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	accountEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	if r.Header.Get("X-Content-Mode") == "partials" {
		if err := accPages.ProfileMain(account(user)).Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accPages.Profile(account(user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) profileHeaderPartial(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	if err := accPages.CommonHeaderOne(
		account(user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) profileLetRecyclePartial(w http.ResponseWriter, r *http.Request) {

	if err := accPages.LetRecycleAndProfile("My Profile").Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) profileChangeDisplayIcon(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	if err := accPages.ProfileChangeDisplayIcon(user.AvatarUrl).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) profileUpdateProfilePartial(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	if err := accPages.ProfileUpdataAccountInfoPartial(account(user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) changeProfileDisplayName(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("err parsing form data; %s", err.Error()))
		return
	}

	name := r.Form.Get("nickname")
	if len(name) > 20 {
		name = name[:19]
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	acc, err := a.service.UpdateAccountDisplayName(
		context.Background(), accountEmail, name)

	if err != nil {
		a.serverError(w, err)
		return
	}

	a.updateAccountCache(context.Background(), acc)

	if err := accPages.CommonHeaderOne(account(acc)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
