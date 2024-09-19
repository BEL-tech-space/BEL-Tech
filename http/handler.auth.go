package http

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func(a *App) signInWithProvider(c echo.Context) error {
	
	var provider string
	if provider = c.Param("provider"); provider == "" {
		return c.String(http.StatusBadRequest, "Provider not specified")
	}

	if provider != "google" && provider != "facebook" {
		return c.String(http.StatusBadRequest, "Invalid provider")
	}

	q := c.Request().URL.Query()
	q.Add("provider", provider)
	c.Request().URL.RawQuery = q.Encode()

	req, res := c.Request(), c.Response().Writer
	
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		return c.JSON(http.StatusOK, gothUser)
	}

	gothic.BeginAuthHandler(res, req)

	return nil
}

func (a *App) callbackHandler(c echo.Context) error {
	var (
		user goth.User
		err  error
	)
	req, res := c.Request(), c.Response().Writer

	if user, err = gothic.CompleteUserAuth(res, req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	val := reflect.ValueOf(user)
	typ := reflect.TypeOf(user)
	
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name
		fieldType := typ.Field(i).Type
		fieldValue := val.Field(i).Interface()

		fmt.Printf("%v:%v:%v\n\n", fieldName, fieldType, fieldValue)

		c.Set("user", user)
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/beltech/dashboard")
}


func(a *App) logout(c echo.Context) error {

	req, res := c.Request(), c.Response().Writer
	gothic.Logout(res, req)
	res.Header().Set("Location", "/index")
	res.WriteHeader(http.StatusTemporaryRedirect)

	return nil
}


// func (a *App) loginUser(c echo.Context) error {
	
// 	var err error
// 	user := c.Get("user").(goth.User) 
	
	
// 	return nil
// }s

 