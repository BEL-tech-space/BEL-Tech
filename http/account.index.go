package http

import (
	accountPages "beltech/frontend/pages/account"
	"context"
	"net/http"
)

func (a *App) homePageHandler(w http.ResponseWriter, r *http.Request) {
	if err := accountPages.IndexPage().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}

}
