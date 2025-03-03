package http

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	accPages "beltech/frontend/pages/account"
	"beltech/service"

	"github.com/go-chi/chi/v5"
)

func (a *App) wasteHistoryHandler(w http.ResponseWriter, r *http.Request) {

	var wasteType string
	if wasteType = chi.URLParam(r, "wasteType"); wasteType == "" {
		http.Error(w, "WasteType not specified", http.StatusBadRequest)
		return
	}

	if wasteType != "glass" && wasteType != "metal" &&
		wasteType != "paper" && wasteType != "plastics" {
		http.Error(w, "Invalid wasteType", http.StatusBadRequest)
		return
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		fmt.Println("error occured here")
		a.serverError(w, err)
		return
	}

	wasteLog, err := a.service.GetWasteLogByAccountPaginated(
		r.Context(),
		accountEmail,
		30,
		0,
	)
	if err != nil {
		a.serverError(w, fmt.Errorf("error getting wastelog: %w", err))
	}

	sortedWaste := sortWasteType(wasteLog)

	var needed []accPages.WasteTypelog
	switch wasteType {
	case "glass":
		needed = sortedWaste.Glass
	case "metal":
		needed = sortedWaste.Metal
	case "plastics":
		needed = sortedWaste.Plastics
	case "paper":
		needed = sortedWaste.Paper
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WastePage(user, wasteType, needed).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sortWasteType(wasteLog []service.WasteLogFromCollector) accPages.WasteSorted {
	glass := make([]accPages.WasteTypelog, 0, len(wasteLog))
	metal := make([]accPages.WasteTypelog, 0, len(wasteLog))
	plastics := make([]accPages.WasteTypelog, 0, len(wasteLog))
	paper := make([]accPages.WasteTypelog, 0, len(wasteLog))

	for _, v := range wasteLog {
		if v.GlassWeighed > 0.000000000 {
			glass = append(glass, accPages.WasteTypelog{
				WasteID:           v.WasteID,
				AgentID:           v.AgentID,
				Claimed:           v.Claimed,
				WasteWeighed:      v.GlassWeighed,
				WastePointsEarned: v.GlassPointsEarned,
				LocationAddr:      v.LocationAddr,
				TxnName:           v.TxnName,
				CreatedAt:         v.CreatedAt,
			})

		}
		if v.MetalWeighed > 0.000000000 {
			metal = append(metal, accPages.WasteTypelog{
				WasteID:           v.WasteID,
				AgentID:           v.AgentID,
				Claimed:           v.Claimed,
				WasteWeighed:      v.MetalWeighed,
				WastePointsEarned: v.MetalPointsEarned,
				LocationAddr:      v.LocationAddr,
				TxnName:           v.TxnName,
				CreatedAt:         v.CreatedAt,
			})
		}

		if v.PlasticsWeighed > 0.000000000 {
			plastics = append(plastics, accPages.WasteTypelog{
				WasteID:           v.WasteID,
				AgentID:           v.AgentID,
				Claimed:           v.Claimed,
				WasteWeighed:      v.PlasticsWeighed,
				WastePointsEarned: v.PlasticsPointsEarned,
				LocationAddr:      v.LocationAddr,
				TxnName:           v.TxnName,
				CreatedAt:         v.CreatedAt,
			})
		}

		if v.PaperWeighed > 0.000000000 {
			paper = append(paper, accPages.WasteTypelog{
				WasteID:           v.WasteID,
				AgentID:           v.AgentID,
				Claimed:           v.Claimed,
				WasteWeighed:      v.PaperWeighed,
				WastePointsEarned: v.PaperPointsEarned,
				LocationAddr:      v.LocationAddr,
				TxnName:           v.TxnName,
				CreatedAt:         v.CreatedAt,
			})
		}
	}
	return accPages.WasteSorted{
		Glass:    slices.Clip(glass),
		Metal:    slices.Clip(metal),
		Plastics: slices.Clip(plastics),
		Paper:    slices.Clip(paper),
	}
}

func (a *App) wasteMainTagPartial(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WasteMainOpeninTag().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) wasteQuickInfo(w http.ResponseWriter, r *http.Request) {

	var wasteType string
	if wasteType = chi.URLParam(r, "wasteType"); wasteType == "" {
		http.Error(w, "WasteType not specified", http.StatusBadRequest)
		return
	}

	if wasteType != "glass" && wasteType != "metal" && wasteType != "paper" && wasteType != "plastics" {
		http.Error(w, "Invalid wasteType", http.StatusBadRequest)
		return
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		fmt.Println("error occured here")
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WasteQuickInfo(wasteType, user).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) wasteFilterOptions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WasteFilterOptionsPartial().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) wasteHistory(w http.ResponseWriter, r *http.Request) {

	var wasteType string
	if wasteType = chi.URLParam(r, "wasteType"); wasteType == "" {
		http.Error(w, "WasteType not specified", http.StatusBadRequest)
		return
	}

	if wasteType != "glass" && wasteType != "metal" && wasteType != "paper" && wasteType != "plastics" {
		http.Error(w, "Invalid wasteType", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WasteHistoryComp(wasteType, nil).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
