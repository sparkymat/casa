package handler

import (
	"net/http"
	"strconv"

	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"github.com/labstack/echo/v4"
)

func CreateHomeItem(_ configiface.ConfigAPI, db dbiface.DatabaseAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := currentUser(c)

		appIDString := c.Param("id")

		appID, err := strconv.ParseUint(appIDString, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid app ID")
		}

		_, err = db.CreateHomeItem(c.Request().Context(), user.ID, uint(appID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete home item")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/apps")
	}
}
