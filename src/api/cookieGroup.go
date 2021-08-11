package api

import (
	"blueBot_go_webserver_echo/src/api/handlers"

	"github.com/labstack/echo/v4"
)

func CookieGroup(cookieGroup *echo.Group) {
	cookieGroup.GET("/main", handlers.MainCookie)
}
