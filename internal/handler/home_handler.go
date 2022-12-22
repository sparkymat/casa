package handler

import (
	"net/http"

	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"git.orion.home/oxhead/casa/view"
	"github.com/labstack/echo/v4"
)

func Home(cfg configiface.ConfigAPI, db dbiface.DatabaseAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := currentUser(c)

		homeItems, err := db.ListHomeItems(c.Request().Context(), user.ID)
		if err != nil {
			//nolint:wrapcheck
			return c.String(http.StatusInternalServerError, "internal server error")
		}

		pageHTML := view.Home(homeItems)
		htmlString := view.Layout(cfg.Title(), pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
