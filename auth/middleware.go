package auth

import (
	"errors"
	"net/http"

	"git.orion.home/oxhead/casa/database/dbiface"
	"git.orion.home/oxhead/casa/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const CurrentUserKey = "current_user"

func Middleware(db dbiface.DatabaseAPI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(CurrentUserKey, &model.User{
				HomeItems: []model.HomeItem{
					{
						CatalogItem: model.CatalogItem{
							URL:      "https://photos.orion.home",
							Title:    "Photos",
							ImageURL: "/uploads/01GMQB5AT691XB5DTJ80GN8KEZ.jpg",
						},
					},
				},
			},
			)
			return next(c)

			email := c.Request().Header.Get("Remote-Email")
			name := c.Request().Header.Get("Remote-Name")

			if email == "" || name == "" {
				//nolint:wrapcheck
				return c.String(http.StatusUnauthorized, "unathorized")
			}

			user, err := db.GetUser(c.Request().Context(), email)
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
				user, err = db.CreateUser(c.Request().Context(), email, name)
			}

			if err != nil {
				//nolint:wrapcheck
				return c.String(http.StatusUnauthorized, "unathorized")
			}

			c.Set(CurrentUserKey, user)

			return next(c)
		}
	}
}
