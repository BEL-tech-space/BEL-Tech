package http

import (
	accountPages "beltech/frontend/pages/account"
	"context"
	"net/http"
	"strings"
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
		if err := accountPages.SpendMainPartial(user.AvatarURL).Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accountPages.SpendPage(user.AvatarURL).Render(context.TODO(), w); err != nil {
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
	if err := accountPages.SpendPageHeader(account(*user).AvatarUrl, accountPages.BeforeHeaderTagPartial(strings.Split(r.URL.Path, "/")[2])).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}