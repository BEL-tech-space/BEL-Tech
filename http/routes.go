package http

import (
	frontend "beltech/frontend"
	"beltech/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (a *App) Routes(cfg utils.EnvConfig) *echo.Echo {

	e := echo.New()
	fs := echo.MustSubFS(frontend.AssetsDir, "assets")
	e.StaticFS("/assets", fs)

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/auth/:provider", a.signInWithProvider)
	e.GET("/auth/:provider/callback", a.callbackHandler)
	e.GET("/logout/:provider", a.logout)

	userProfile := e.Group("/beltech")
	// userProfile.Use(session.Middleware(sessions.NewCookieStore([]byte(cfg.SessionSecret))))
	userProfile.GET("/dashboard", a.homePageHandler)
	userProfile.GET("/scan", a.scanQRPageHandler)
	userProfile.GET("/bins", a.locationPageHandler)
	userProfile.GET("/wallet", a.walletPageHandler)
	userProfile.GET("/exchange", a.exchangePageHandler)

	e.GET("/signup", a.signupPageHandler)
	e.GET("/signin", a.signinPageHandler)
	e.GET("/verify", a.verifyPageHandler)
	e.GET("/reset", a.resetPageHandler)
	e.GET("/verify_success", a.verify_successPageHandler)
	e.GET("/password_reset", a.password_resetPageHandler)
	e.GET("/index", a.splash_screenPageHandler)

	return e
}
