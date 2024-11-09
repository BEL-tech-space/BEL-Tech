package http

import (
	agentPages "beltech/frontend/pages/agent"
	"context"
	"net/http"
)

func (a *App) agentDashboard(w http.ResponseWriter, r *http.Request) {
	if err := agentPages.DashboardPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) historyPage(w http.ResponseWriter, r *http.Request) {
	if err := agentPages.HistoryPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) profilePage(w http.ResponseWriter, r *http.Request) {
	if err := agentPages.ProfilePage(
		agentPages.Collector{
			Email:     "giwaoluatobi@gmail.com",
			AvatarUrl: "https://lh3.googleusercontent.com/a-/ALV-UjWhunlRgNtwAAdFNTue7WN3k01azp8GY0qJHQqsdfXBB26X40qIIw=s96-c",
			NickName:  "Tobias",
			Role:      "Collector",
		},
	).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
	if err := agentPages.UnauthorizedPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
