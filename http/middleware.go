package http

import (
	"beltech/service"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
)

func (a *App) logInUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			user service.User
			err  error
		)

		u := c.Get("user").(goth.User)

		if user, err = mapGothUserToServiceUser(u); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		fmt.Println(user)

		// if err = a.service.CreateOrUpdateUser(context.Background(), user); err != nil {
		// 	return err
		// }

		if err = createNewUserSession(c, user); err != nil {
			return err
		}

		return next(c)
	}
}

func createNewUserSession(c echo.Context, user service.User) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
	}

	user.IsAuthenticated = true
	sess.Values[user.UserID] = user

	return sess.Save(c.Request(), c.Response())
}



func mapGothUserToServiceUser(gothUser goth.User) (service.User, error) {
	user := service.User{}
	user.UserID = gothUser.UserID
	user.Email = gothUser.Email
	user.FirstName = pgtype.Text{String: gothUser.FirstName, Valid: true}
	user.LastName = pgtype.Text{String: gothUser.LastName, Valid: true}
	user.NickName = pgtype.Text{String: gothUser.NickName, Valid: true}
	user.AvatarUrl = pgtype.Text{String: gothUser.AvatarURL, Valid: true}
	user.AccessToken = pgtype.Text{String: gothUser.AccessToken, Valid: true}
	user.AccessTokenExpires = pgtype.Timestamp{Time: gothUser.ExpiresAt, Valid: true}
	user.IDToken = pgtype.Text{String: gothUser.IDToken, Valid: true}
	return user, nil
}

