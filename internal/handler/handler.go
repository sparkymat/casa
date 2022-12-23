package handler

import (
	"git.orion.home/oxhead/casa/auth"
	"git.orion.home/oxhead/casa/model"
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
