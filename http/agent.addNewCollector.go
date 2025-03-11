package http

import (
	agentPages "beltech/frontend/pages/agent"
	"beltech/service"
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (a *App) addNewCollector(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value(contextKeyRole).(service.Collectorrole) != service.CollectorroleSuperCollector {
		a.clientError(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	collectorEmail := r.Context().Value(contextKeyEmail).(string)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.AddNewCollector(a.collectors[collectorEmail], false).Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
func (a *App) addNewCollectorPost(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value(contextKeyRole).(service.Collectorrole) != service.CollectorroleSuperCollector {
		a.clientError(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	err := r.ParseForm()
	if err != nil {
		a.clientError(w, http.StatusBadRequest, err)
		return
	}

	displayName := r.PostForm.Get("displayname")
	firstName := r.PostForm.Get("firstname")
	lastName := r.PostForm.Get("lastname")
	email := r.PostForm.Get("email")
	role := r.PostForm.Get("role")

	fmt.Println("role")

	if email == "" || firstName == "" || lastName == "" || role == "" {
		a.clientError(w, http.StatusBadRequest, errors.New("'email, 'firstname', 'lastname' or 'role' cannot be empty"))
		return
	}

	var collector service.Collector

	if role == "super_collector" {
		collector, err = a.service.AddNewAgentAsSuperCollector(
			context.Background(),
			service.AddNewAgentAsSuperCollectorParams{
				CollectorEmail: email,
				FirstName:      firstName,
				LastName:       lastName,
				DisplayName:    displayName,
			},
		)

	} else {
		collector, err = a.service.AddNewAgentAsCollector(
			context.Background(),
			service.AddNewAgentAsCollectorParams{
				CollectorEmail: email,
				FirstName:      firstName,
				LastName:       lastName,
				DisplayName:    displayName,
			},
		)
	}

	if err != nil {
		a.serverError(w, err)
	}

	a.collectors[collector.CollectorEmail] = collector

	collectorEmail := r.Context().Value(contextKeyEmail).(string)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.AddNewCollector(a.collectors[collectorEmail], true).Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}
