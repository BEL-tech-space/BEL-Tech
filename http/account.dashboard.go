package http

import (
	accPages "beltech/frontend/pages/account"
	"beltech/service"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (a *App) dashboardPageHandler(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	isNewAccount := false
	if cookie, err := r.Cookie("isNewAccount"); err == nil && cookie.Value == "true" {
		http.SetCookie(w, &http.Cookie{
			Name:   "isNewAccount",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})

		isNewAccount = true
	}

	if isNewAccount {
		go func() {
			err := a.sendGreetingEmail(accountEmail)
			if err != nil {
				fmt.Println("error sending email:", err)
			}
		}()
	}

	res, err := a.service.GetWasteByAccountForMonth(
		r.Context(),
		accountEmail,
		deduceTime(0),
	)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	chart, isEmpty := produceChartHTML(sumWasteLogForGraph(res), deduceTime(0))
	if err := accPages.DashboardPage(user, chart, isEmpty, isNewAccount).
		Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardMainTagPartial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardMainOpeningTag().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardSwiperOne(w http.ResponseWriter, r *http.Request) {
	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardSwiperOne(user).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardDataViz(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	res, err := a.service.GetWasteByAccountForMonth(
		r.Context(),
		accountEmail,
		deduceTime(0),
	)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardDataViz(
		produceChartHTML(sumWasteLogForGraph(res), deduceTime(0))).
		Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func sumWasteLogForGraph(w []service.WasteLogFromCollector) [4]float64 {
	var glass, metal, plastics, paper float64
	for _, v := range w {
		glass += float64(v.GlassWeighed)
		metal += float64(v.MetalWeighed)
		plastics += float64(v.PlasticsWeighed)
		paper += float64(v.PaperWeighed)
	}

	return [4]float64{glass, metal, plastics, paper}
}

func (a *App) dashboardSwiperTwo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardSwiperTwo().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardEchart(w http.ResponseWriter, r *http.Request) {
	month := r.URL.Query().Get("month")

	num, err := strconv.Atoi(month)
	if err != nil {
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("invalid query parameter for month: %s", err.Error()))
		return
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	res, err := a.service.GetWasteByAccountForMonth(
		r.Context(),
		accountEmail,
		deduceTime(num),
	)
	if err != nil {
		a.serverError(w, err)
		return
	}

	if err := accPages.DashboardEchart(
		produceChartHTML(sumWasteLogForGraph(res), deduceTime(num))).
		Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func deduceTime(v int) time.Time {
	if v > 4 {
		v = 4
	}

	return time.Now().AddDate(0, -v, 0)
}
