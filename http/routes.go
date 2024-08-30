package http

import (
	"github.com/labstack/echo/v4"
)


func (a *App) Routes() *echo.Echo {
	e := echo.New()

	e.POST("/addWasteBin", a.addWasteBin)


	return e
}