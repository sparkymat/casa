package route

import (
	"fmt"
	"net/http"
	"path"

	"git.orion.home/oxhead/casa/auth"
	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"git.orion.home/oxhead/casa/internal/handler"
	"git.orion.home/oxhead/casa/view"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func customErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := err.Error()

	//nolint:errorlint
	if httpErr, ok := err.(*echo.HTTPError); ok {
		code = httpErr.Code
		message = fmt.Sprintf("%v", httpErr.Message)
	}

	c.Logger().Error(err)

	pageHTML := view.Error(fmt.Sprintf("%v", message))
	htmlString := view.BasicLayout("error", pageHTML)

	if err := c.HTMLBlob(code, []byte(htmlString)); err != nil {
		c.Logger().Error(err)
	}
}

func Setup(e *echo.Echo, cfg configiface.ConfigAPI, db dbiface.DatabaseAPI) {
	e.HTTPErrorHandler = customErrorHandler
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	e.Static("/css", "public/css")
	e.Static("/js", "public/js")
	e.Static("/fonts", "public/fonts")
	e.Static("/uploads", path.Join(cfg.DataPath(), "uploads"))

	app := e.Group("")
	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	app.Use(auth.Middleware(db))

	app.GET("/", handler.Home(cfg, db))
	app.GET("/apps", handler.Apps(cfg, db))
	app.GET("/apps/new", handler.NewApp(cfg, db))
	app.POST("/apps", handler.CreateApp(cfg, db))
	app.POST("/apps/:id/delete", handler.DestroyApp(cfg, db))
	app.POST("/apps/:id/home_items", handler.CreateHomeItem(cfg, db))
	app.POST("/apps/:id/home_items/:home_item_id/delete", handler.DestroyHomeItem(cfg, db))
}
