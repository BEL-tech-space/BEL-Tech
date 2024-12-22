package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"fmt"
	"net/http"
)

func (a *App) airtimePageHandler(w http.ResponseWriter, r *http.Request) {
	accountEmail := r.Context().Value(contextKeyEmail).(string)

	if err := r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, err)
		return
	}

	for v, c := range r.PostForm {
		fmt.Println(v, c)
	}

	user, err := a.userFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Header.Get("X-Content-Mode") == "partials" {
		if err := accPages.AirtimeMain().Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accPages.AirtimePage(user.AvatarURL).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) airtimePageHeaderPartial(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.AirtimeHeaderPartialComp(
		account(*user).AvatarUrl,
		accPages.BeforeHeaderTagPartial(
			"Airtime"),
		accPages.AirtimeMainOpeningTagComp()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) airtimeInputNumberPartial(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.AirtimeInputNumberPartialcomp().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) airtimeBeneficiaries(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.AirtimeBeneficiariesComp().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
