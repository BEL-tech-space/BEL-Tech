package http

import (
	accPages "beltech/frontend/pages/account"
	agentPages "beltech/frontend/pages/agent"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (a *App) agentWasteHistoryHandler(w http.ResponseWriter, r *http.Request) {

	var wasteType string
	if wasteType = chi.URLParam(r, "wasteType"); wasteType == "" {
		a.clientError(w, http.StatusBadRequest, errors.New("wasteType not specified"))
		return
	}

	if wasteType != "glass" && wasteType != "metal" &&
		wasteType != "paper" && wasteType != "plastics" {
		a.clientError(w, http.StatusBadRequest, errors.New("invalid wasteType"))
		return
	}

	agentId := r.Context().Value(contextKeyAgentID).(uuid.UUID)
	wasteLog, err := a.service.GetWasteLogByCollectorPaginated(
		context.Background(),
		agentId,
		30,
		0,
	)
	if err != nil {
		a.serverError(w, fmt.Errorf("error getting wastelog: %w", err))
		return
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

	collectorEmail := r.Context().Value(contextKeyEmail).(string)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.WastePage(a.collectors[collectorEmail], wasteType, needed).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
