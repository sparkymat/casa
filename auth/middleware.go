package auth

import (
	"errors"
	"net/http"

	"git.orion.home/oxhead/casa/database/dbiface"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const CurrentUserKey = "current_user"

func Middleware(db dbiface.DatabaseAPI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			email := c.Request().Header.Get("Remote-Email")
			name := c.Request().Header.Get("Remote-Name")

			if email == "" || name == "" {
				return c.String(http.StatusUnauthorized, "unathorized")
			}

			user, err := db.GetUser(c.Request().Context(), email)
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
				user, err = db.CreateUser(c.Request().Context(), email, name)
			}
			if err != nil {
				return c.String(http.StatusUnauthorized, "unathorized")
			}

			c.Set(CurrentUserKey, user)

			return next(c)
		}
	}
}
