package http

import (
	"context"
	"fmt"
	"net/http"

	accPages "beltech/frontend/pages/account"

	"github.com/a-h/templ"
)

func (a *App) beneficiariesPageHandler(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		fmt.Println("error occured here")
		a.serverError(w, err)
		return
	}

	if err := accPages.BeneficiariesPage(user.AvatarURL).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) beneficiariesPageHeaderPartial(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.BeneficiariesHeaderPartialComp(
		account(*user).AvatarUrl,
		accPages.BeforeHeaderTagPartial(
			"Spend"),
		templ.NopComponent).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) beneficiariesMain(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.BeneficiariesMain().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
