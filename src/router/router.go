package router

import (
	"blueBot_go_webserver_echo/src/api"
	"blueBot_go_webserver_echo/src/api/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	echoInstance := echo.New()

	// Create groups:
	adminGroup := echoInstance.Group("/admin")
	// Cookies
	cookieGroup := echoInstance.Group("/cookie")
	// JWT:
	jwtGroup := echoInstance.Group("/jwt")

	// Set all middleware
	middleware.SetMainMiddleware(echoInstance)
	middleware.SetAdminMiddleware(adminGroup)
	middleware.SetCookieMiddleware(cookieGroup)
	middleware.SetJwtMiddleware(jwtGroup)

	// Main routes
	api.MainGroup(echoInstance)

	// Group routes
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	return echoInstance
}
