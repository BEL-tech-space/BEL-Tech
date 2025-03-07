package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"errors"
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
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("failed to parse form: error: %w", err))
		return
	}

	wasteUUIDstr := r.PostFormValue("uuid")
	waste, ok := a.wasteLogCache.Get(wasteUUIDstr)
	if !ok {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := accPages.ErorOnSubmitWaste("waste log does not exist").Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if waste.Claimed {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := accPages.ErorOnSubmitWaste("waste already claimed").Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if a.config.Build != productionEnv {
		time.Sleep(4 * time.Second)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.SubmitWaste().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) submitWaste(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, errors.New("failed to parse form"))
		return
	}

	wasteUUIDstr := r.PostFormValue("uuid")
	waste, ok := a.wasteLogCache.Get(wasteUUIDstr)
	if !ok {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := accPages.ErorOnSubmitWaste("waste log does not exist").Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if waste.Claimed {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := accPages.ErorOnSubmitWaste("waste already claimed").Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
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
		context.Background(), waste); err != nil {
		a.serverError(w, err)
		return
	}

	if err := a.addUserToCache(context.Background(), accountEmail); err != nil {
		// a.serverError(w, err)
	}

	if err := a.updateCollectorInfo(waste.CollectorEmail); err != nil {
		// a.serverError(w, err)
		// return
	}

	if a.config.Build != productionEnv {
		time.Sleep(4 * time.Second)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.ClaimSuccessful(waste.TotalPointsEarned).Render(r.Context(), w); err != nil {
		a.serverError(w, err)
	}
}
