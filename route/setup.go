package route

import (
	"net/http"

	"git.orion.home/oxhead/casa/auth"
	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo, _ configiface.ConfigAPI, db dbiface.DatabaseAPI) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(auth.Middleware(db))

	e.GET("/", func(c echo.Context) error {
		//nolint:wrapcheck
		return c.String(http.StatusOK, "hello, world")
	})
}
