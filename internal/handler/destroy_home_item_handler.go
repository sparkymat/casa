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
			//nolint:wrapcheck
			return c.String(http.StatusUnprocessableEntity, "invalid app")
		}

		err = db.DestroyHomeItem(c.Request().Context(), uint(appID))
		if err != nil {
			//nolint:wrapcheck
			return c.String(http.StatusInternalServerError, "app delete failed")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/apps")
	}
}
