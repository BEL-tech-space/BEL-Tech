package http

import (
	"beltech/service"
	"context"
	"errors"
	"net/http"
)

func (a *App) addNewCollectorPost(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value(contextKeyRole) != "super_collector" {
		// not authorize page
		return
	}

	err := r.ParseForm()
	if err != nil {
		a.clientError(w, http.StatusBadRequest, err)
		return
	}

	nickName := r.PostForm.Get("nickName")
	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")

	email := r.PostForm.Get("email")
	role := r.PostForm.Get("role")
	is := r.PostForm.Get("is")

	if email == "" || role == "" || is == "" {
		a.clientError(w, http.StatusBadRequest, errors.New("'email, 'role' or 'is' cannot be empty"))
		return
	}

	agentID, err := a.service.AddNewAgentAndCollector(context.Background(), service.AddNewAgentParams{
		Is:        is,
		Email:     email,
		Role:      role,
		FirstName: firstName,
		LastName:  lastName,
		NickName:  nickName,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.collectors = append(a.collectors, collector{Email: "", AgentID: agentID, Role: ""})

}
