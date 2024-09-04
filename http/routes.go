package http

import (
	frontend "beltech/frontend"

	"github.com/labstack/echo/v4"
)


func (a *App) Routes() *echo.Echo {
	

	e := echo.New()
	fs := echo.MustSubFS(frontend.AssetsDir, "assets")
	e.StaticFS("/assets", fs)

	e.GET("/home", a.homePageHandler)
	e.GET("/scan", a.scanQRPageHandler)
	e.GET("/location", a.locationPageHandler)
	e.GET("/wallet", a.walletPageHandler)
	e.GET("/exchange", a.exchangePageHandler)
	
	e.POST("/addWasteBin", a.addWasteBin)

	return e
}