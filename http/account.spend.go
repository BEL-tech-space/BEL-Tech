package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) spendPageHandler(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Header.Get("X-Content-Mode") == "partials" {
		if err := accPages.SpendMain(account(*user)).Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accPages.SpendPage(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) spendPageHeaderPartial(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.SpendHeaderPartialComp(
		account(*user).AvatarUrl,
		accPages.BeforeHeaderTagPartial(
			"Spend BEL Token"),
		accPages.SpendMainOpeningTagComp()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) spendBalance(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.SpendBalanceComp(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) spendOtherHalfPartial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.SpendOtherHalfPartialComp().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
