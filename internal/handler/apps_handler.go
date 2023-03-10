package handler

import (
	"net/http"

	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"git.orion.home/oxhead/casa/presenter"
	"git.orion.home/oxhead/casa/view"
	"github.com/labstack/echo/v4"
)

func Apps(cfg configiface.ConfigAPI, db dbiface.DatabaseAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := currentUser(c)

		csrfToken := getCSRFToken(c)
		if csrfToken == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Form verification failed")
		}

		catalogItems, err := db.ListCatalogItems(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Database read failed")
		}

		homeItems, err := db.ListHomeItems(c.Request().Context(), user.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Database read failed")
		}

		appItems := presenter.AppItemsFromModelList(catalogItems, homeItems)

		pageHTML := view.Apps(csrfToken, appItems)
		htmlString := view.Layout(cfg.Title(), pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
