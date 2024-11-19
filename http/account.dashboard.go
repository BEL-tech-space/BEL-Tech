package http

import (
	accountPages "beltech/frontend/pages/account"
	"context"
	"fmt"
	"net/http"

	"math/rand"

	"github.com/go-chi/chi/v5"
)

func (a *App) accountDashboard(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		fmt.Println("error occured here")
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.AccountDashboardPage(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) accountProfilePage(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.AccountProfile(account(*user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) rechargeAirtime(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(contextKeyEmail).(string)

	if err := r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, err)
		return
	}

	for v, c := range r.PostForm {
		fmt.Println(v, c)
	}

	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	if err := accountPages.RechargeAirtimePage(account(*user)).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) airtimeBeneficiaries(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		fmt.Println("error occured here")
		a.serverError(w, err)
		return
	}

	if err := accountPages.AirtimeBeneficiaryPage(account(*user)).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) wasteHistoryHandler(w http.ResponseWriter, r *http.Request) {

	var wasteType string
	if wasteType = chi.URLParam(r, "wasteType"); wasteType == "" {
		http.Error(w, "WasteType not specified", http.StatusBadRequest)
		return
	}

	if wasteType != "glass" && wasteType != "metal" && wasteType != "paper" && wasteType != "plastics" {
		http.Error(w, "Invalid wasteType", http.StatusBadRequest)
		return
	}

	userEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		fmt.Println("error occured here")
		a.serverError(w, err)
		return
	}

	if err := accountPages.WasteHistoryPage(account(*user), wasteType, randomArray()).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) wasteTipsHandler(w http.ResponseWriter, r *http.Request) {

	var wasteType string
	if wasteType = chi.URLParam(r, "wasteType"); wasteType == "" {
		http.Error(w, "WasteType not specified", http.StatusBadRequest)
		return
	}

	if wasteType != "glass" && wasteType != "metal" && wasteType != "paper" && wasteType != "plastics" {
		http.Error(w, "Invalid wasteType", http.StatusBadRequest)
		return
	}

	if err := accountPages.TipsOnRecyclingPage(wasteType).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func randomArray() []string {
	if rand.Intn(2) == 0 {
		return []string{} 
	}
	return []string{"example1", "example2", "example3"}
}

func (a *App) allWasteHistoryHandler(w http.ResponseWriter, r *http.Request) {


	if err := accountPages.AllWasteHistoryPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

