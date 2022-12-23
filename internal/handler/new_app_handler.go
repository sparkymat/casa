package handler

import (
	"net/http"

	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"git.orion.home/oxhead/casa/view"
	"github.com/labstack/echo/v4"
)

func NewApp(cfg configiface.ConfigAPI, db dbiface.DatabaseAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := getCSRFToken(c)
		if csrfToken == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Form verification failed")
		}

		categories, err := db.ListCategories(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load categories")
		}

		pageHTML := view.NewApp(csrfToken, categories)
		htmlString := view.Layout(cfg.Title(), pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
