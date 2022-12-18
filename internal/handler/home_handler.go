package handler

import (
	"net/http"

	"git.orion.home/oxhead/casa/view"
	"github.com/labstack/echo/v4"
)

func Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		pageHTML := view.Home()
		htmlString := view.Layout("archmark", pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
