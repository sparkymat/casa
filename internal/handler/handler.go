package handler

import (
	"net/http"

	"git.orion.home/oxhead/casa/auth"
	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/model"
	"git.orion.home/oxhead/casa/view"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getCSRFToken(c echo.Context) string {
	csrfTokenVal := c.Get(middleware.DefaultCSRFConfig.ContextKey)
	if csrfTokenVal == nil {
		return ""
	}

	csrfToken, ok := csrfTokenVal.(string)
	if !ok {
		return ""
	}

	return csrfToken
}

func currentUser(c echo.Context) *model.User {
	user, isUser := c.Get(auth.CurrentUserKey).(*model.User)
	if !isUser {
		return nil
	}

	return user
}

func renderError(c echo.Context, cfg configiface.ConfigAPI, errorMessage string) error {
	pageHTML := view.Error(errorMessage)
	htmlString := view.Layout(cfg.Title(), pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
