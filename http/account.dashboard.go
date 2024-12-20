package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
	"strings"
)

func (a *App) dashboardPageHandler(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Header.Get("X-Content-Mode") == "partials" {
		if err := accPages.DashboardMain(chart()).Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := accPages.DashboardPage(account(*user), chart()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardHeaderPartial(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := accPages.CommonHeaderOne(
		account(*user),
		accPages.BeforeHeaderTagPartial(strings.Split(r.URL.Path, "/")[2]),
		accPages.DashboardMainOpeningTagComp()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardLetRecyclePartial(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.LetRecycleAndProfileComp("Let's Recycle!").Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardSwiperOne(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardSwiperOneComp().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardRecycledAndSeeAllPartial(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardRecycledAndSeeAllComp().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardDataViz(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardDataVizComp(chart()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardTipsParagraph(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardTipsParagraphComp().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardSwiperTwo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardSwiperTwoComp().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
