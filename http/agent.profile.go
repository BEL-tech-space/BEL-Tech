package http

import (
	agentPages "beltech/frontend/pages/agent"
	"context"
	"net/http"
)

func (a *App) profilePage(w http.ResponseWriter, r *http.Request) {

	collectorEmail := r.Context().Value(contextKeyEmail).(string)
	collector, ok := a.collectors[collectorEmail]
	if !ok {
		http.Redirect(w, r, "/agent/login.html", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.ProfilePage(collector).Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) upUpdateCollectorDisplayName(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, err)
		return
	}

	name := r.Form.Get("nickName")
	if len(name) > 20 {
		name = name[:19]
	}

	collectorEmail := r.Context().Value(contextKeyEmail).(string)

	collector, err := a.service.UpdateCollectorDisplayName(
		context.Background(),
		collectorEmail,
		name,
	)

	if err == nil {
		a.collectors[collectorEmail] = collector
	}

	if err != nil {
		a.serverError(w, err)
	}

	http.Redirect(w, r, "/agent/profile", http.StatusSeeOther)
}
