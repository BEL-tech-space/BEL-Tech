package http

import (
	agentPages "beltech/frontend/pages/agent"
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (a *App) agentDashboard(w http.ResponseWriter, r *http.Request) {

	agentId := r.Context().Value(contextKeyAgentID).(uuid.UUID)
	res, err := a.service.GetWasteByCollectorForMonth(
		context.Background(),
		agentId,
		deduceTime(0),
	)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.DashboardPage(res).Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) logInCollector(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.LogInCollector().Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) logOutCollector(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.LogOutCollector().Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) unauthorized(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.UnauthorizedPage().Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
