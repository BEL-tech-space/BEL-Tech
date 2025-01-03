package http

import (
	accPages "beltech/frontend/pages/account"
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
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardPage(account(user), produceChartHTML([4]float64{9, 5, 7, 15}, deduceTime(0))).
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardSwiperOne().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) dashboardDataViz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.DashboardDataViz(
		produceChartHTML([4]float64{9, 5, 7, 15}, deduceTime(4))).
		Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
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

	arr := [4]float64{9, 5, 7, 15}
	switch num {
	case 1:
		arr = [4]float64{23, 15, 20, 9}
	case 2:
		arr = [4]float64{10, 51, 67, 80}
	case 3:
		arr = [4]float64{100, 25, 37, 215}
	}

	if err := accPages.DashboardEchart(
		produceChartHTML(arr,
			deduceTime(num),
		)).Render(context.TODO(), w); err != nil {
	}
}

func deduceTime(v int) time.Time {
	if v > 4 {
		v = 4
	}
	return time.Now().AddDate(0, -v, 0)
}
