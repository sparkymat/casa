package route

import (
	"git.orion.home/oxhead/casa/auth"
	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"git.orion.home/oxhead/casa/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo, cfg configiface.ConfigAPI, db dbiface.DatabaseAPI) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(auth.Middleware(db))

	e.Static("/css", "public/css")
	e.Static("/js", "public/js")

	e.GET("/", handler.Home(cfg))
	e.GET("/apps", handler.Apps(cfg))
}
