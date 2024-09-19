package http

import (
	frontend "beltech/frontend/pages"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *App) signupPageHandler(c echo.Context) error {

	if err := frontend.SignUpPage().Render(context.TODO(), c.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
