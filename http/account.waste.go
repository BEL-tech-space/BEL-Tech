package http

import (
	"context"
	"fmt"
	"net/http"

	accPages "beltech/frontend/pages/account"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/rand"
)

func (a *App) wasteHistoryHandler(w http.ResponseWriter, r *http.Request) {

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

	if err := accPages.WastePage(account(user), wasteType, nil).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func randomArray() []string {
	if rand.Intn(2) == 0 {
		return []string{}
	}
	return []string{"example1", "example2", "example3"}
}
