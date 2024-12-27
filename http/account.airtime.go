package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"fmt"
	"net/http"
)

func (a *App) airtimePageHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, err)
		return
	}

	for v, c := range r.PostForm {
		fmt.Println(v, c)
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	result, err := a.service.GetBeneficiaries(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	if r.Header.Get("X-Content-Mode") == "partials" {
		if err := accPages.AirtimeMain(result).Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accPages.AirtimePage(user.AvatarUrl, result).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) airtimePageHeaderPartial(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	if err := accPages.AirtimeHeaderPartial(
		account(user).AvatarUrl,
		accPages.BeforeHeaderTagPartial(
			"Airtime"),
		accPages.AirtimeMainOpeningTag()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) airtimeInputNumberPartial(w http.ResponseWriter, r *http.Request) {

	if err := accPages.AirtimeInputNumberPartial().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) airtimeBeneficiaries(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)

	result, err := a.service.GetBeneficiaries(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}
	if err := accPages.AirtimeBeneficiaries(result).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
