package handler

import (
	"net/http"
	"strconv"

	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"github.com/labstack/echo/v4"
)

func DestroyApp(_ configiface.ConfigAPI, db dbiface.DatabaseAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		appIDString := c.Param("id")

		appID, err := strconv.ParseUint(appIDString, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid app ID")
		}

		err = db.DestroyCatalogItem(c.Request().Context(), uint(appID))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Failed to delete app")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/apps")
	}
}
