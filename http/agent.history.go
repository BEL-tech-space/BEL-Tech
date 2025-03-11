package http

import (
	agentPages "beltech/frontend/pages/agent"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (a *App) historyPage(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.HistoryPage(wasteLog).Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
