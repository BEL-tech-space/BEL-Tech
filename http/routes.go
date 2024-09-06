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
	e.GET("/scan", a.scanQRPageHandler)
	e.GET("/location", a.locationPageHandler)
	e.GET("/wallet", a.walletPageHandler)
	e.GET("/exchange", a.exchangePageHandler)
	e.GET("/signup", a.signupPageHandler)
	e.GET("/signin", a.signinPageHandler)
	e.GET("/verify", a.verifyPageHandler)
	e.GET("/reset", a.resetPageHandler)
	e.GET("/verify_success", a.verify_successPageHandler)
	e.GET("/password_reset", a.password_resetPageHandler)
	e.GET("/index", a.splash_screenPageHandler)
	

	return e
}
