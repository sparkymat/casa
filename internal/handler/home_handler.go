package handler

import (
	"net/http"

	"git.orion.home/oxhead/casa/auth"
	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/model"
	"git.orion.home/oxhead/casa/view"
	"github.com/labstack/echo/v4"
)

func Home(cfg configiface.ConfigAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get(auth.CurrentUserKey).(*model.User)
		pageHTML := view.Home(user.HomeItems)
		htmlString := view.Layout(cfg.Title(), pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
