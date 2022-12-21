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
			//nolint:wrapcheck
			return c.String(http.StatusInternalServerError, "csrf error")
		}

		catalogItems, err := db.ListCatalogItems(c.Request().Context())
		if err != nil {
			//nolint:wrapcheck
			return c.String(http.StatusInternalServerError, "internal server error")
		}

		appItems := presenter.AppItemsFromModelList(catalogItems, user.HomeItems)

		pageHTML := view.Apps(csrfToken, appItems)
		htmlString := view.Layout(cfg.Title(), pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
