package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"fmt"
	"net/http"
	"time"
)

func (a *App) scanQRPageHandler(w http.ResponseWriter, r *http.Request) {
	accountEmail := r.Context().Value(contextKeyEmail).(string)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.ScanQRPage(accountEmail).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) confirmWaste(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	wasteUUIDstr := r.PostFormValue("uuid")
	glass, paper, metal, plastics, err := getFormattedData(r)
	if err != nil {
		http.Error(w, "Failed to process form", http.StatusBadRequest)
		return
	}

	waste, ok := a.wasteLogCache.Get(wasteUUIDstr)
	if !ok {
		fmt.Println("waste log does not exist")

		// return
	}

	if waste.Claimed {
		fmt.Println("it has been claimed")

		// return
	}

	fmt.Printf("\n found:%v value:%+v\n\n", ok, waste)
	fmt.Println(waste.WasteID.String() == wasteUUIDstr)
	fmt.Println(waste.GlassWeighed == glass)
	fmt.Println(waste.PlasticsWeighed == plastics)
	fmt.Println(waste.PaperWeighed == paper)
	fmt.Println(waste.MetalWeighed == metal)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.SubmitWaste().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) submitWaste(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	wasteUUIDstr := r.PostFormValue("uuid")
	waste, ok := a.wasteLogCache.Get(wasteUUIDstr)
	if !ok {
		fmt.Println("waste log does not exist")

		// return
	}

	if waste.Claimed {
		fmt.Println("it has been claimed")

		// return
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	waste.Claimed = true
	waste.AccountEmail = accountEmail

	if saved := a.wasteLogCache.SetWithTTL(
		wasteUUIDstr,
		waste,
		1,
		24*time.Hour,
	); !saved {
		fmt.Println("claimed waste not saved back to cache")
	}

	if err := a.service.AddNewWasteFromCollector(
		r.Context(), waste); err != nil {
		a.serverError(w, err)
		return
	}

	fmt.Println("ok")

}
