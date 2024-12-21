package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) homePageHandler(w http.ResponseWriter, r *http.Request) {
	if err := accPages.IndexPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}

}
