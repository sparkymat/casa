package route

import (
	"path"

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
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	e.Use(auth.Middleware(db))

	e.Static("/css", "public/css")
	e.Static("/js", "public/js")
	e.Static("/fonts", "public/fonts")
	e.Static("/uploads", path.Join(cfg.DataPath(), "uploads"))

	e.GET("/", handler.Home(cfg))
	e.GET("/apps", handler.Apps(cfg, db))
	e.GET("/apps/new", handler.NewApp(cfg))
	e.POST("/apps", handler.CreateApp(cfg, db))
	e.POST("/apps/:id/delete", handler.DestroyApp(cfg, db))
	e.POST("/apps/:id/home_items", handler.CreateHomeItem(cfg, db))
}
