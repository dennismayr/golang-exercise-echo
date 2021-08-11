package api

import (
	"blueBot_go_webserver_echo/src/api/handlers"

	"github.com/labstack/echo/v4"
)

func JwtGroup(jwtGroup *echo.Group) {
	jwtGroup.GET("/main", handlers.MainJwt)
}
