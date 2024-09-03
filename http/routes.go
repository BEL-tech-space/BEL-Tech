package http

import (
	frontend "beltech/frontend"

	"github.com/labstack/echo/v4"
)


func (a *App) Routes() *echo.Echo {
	

	e := echo.New()
	fs := echo.MustSubFS(frontend.AssetsDir, "assets")
	e.StaticFS("/assets", fs)

	e.POST("/addWasteBin", a.addWasteBin)
	e.GET("/exchange", a.exchangePageHandler)
	e.GET("/scan", a.scanQRPageHandler)
	e.GET("/wallet", a.walletPageHandler)

	return e
}