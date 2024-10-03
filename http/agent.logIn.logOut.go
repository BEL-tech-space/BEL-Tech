package http

import (
	agentPages "beltech/frontend/pages/agent"
	"context"
	"net/http"
)

func (a *App) logInCollector(w http.ResponseWriter, r *http.Request) {
	if err := agentPages.LogInCollector().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) logOutCollector(w http.ResponseWriter, r *http.Request) {
	if err := agentPages.LogOutCollector().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (a *App) unauthorized(w http.ResponseWriter, r *http.Request) {
	if err := agentPages.LogOutCollector().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
