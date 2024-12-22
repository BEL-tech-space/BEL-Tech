package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
	"strings"
)

func (a *App) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	accountEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.userFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Header.Get("X-Content-Mode") == "partials" {
		if err := accPages.ProfileMain(account(*user)).Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accPages.Profile(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) profileHeaderPartial(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := accPages.CommonHeaderOne(
		account(*user),
		accPages.BeforeHeaderTagPartial(strings.Split(r.URL.Path, "/")[2]),
		accPages.ProfileMainOpeningTagComp()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) profileLetRecyclePartial(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.LetRecycleAndProfileComp("My Profile").Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) profileChangeDisplayIcon(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.ProfileChangeDisplayIconComp(user.AvatarURL).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) profileUpdateProfilePartial(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.ProfileUpdataAccountInfoPartialComp(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
