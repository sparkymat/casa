package handler

import (
	"net/http"
	"strconv"

	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"github.com/labstack/echo/v4"
)

func DestroyHomeItem(_ configiface.ConfigAPI, db dbiface.DatabaseAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		appIDString := c.Param("home_item_id")

		appID, err := strconv.ParseUint(appIDString, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid home item ID")
		}

		err = db.DestroyHomeItem(c.Request().Context(), uint(appID))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Failed to delete home item")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/apps")
	}
}
