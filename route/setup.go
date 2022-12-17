package route

import (
	"net/http"

	"git.orion.home/oxhead/casa/config/configiface"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo, _ configiface.ConfigAPI) {
	e.GET("/", func(c echo.Context) error {
		//nolint:wrapcheck
		return c.String(http.StatusOK, "hello, world")
	})
}
