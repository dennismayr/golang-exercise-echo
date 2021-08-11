package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Main
func SetMainMiddleware(echoInstance *echo.Echo) {
	echoInstance.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "./static",
	}))
	echoInstance.Use(serverHeader)
}

// Headers
// Funcs that start in lowercase are PRIVATE
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	// HandlerFunc defines a function to serve HTTP requests
	return func(c echo.Context) error {
		// c.Response().Header().Set(echo.HeaderServer, "NinoHeaders/1.0")
		c.Response().Header().Set("SuperHeader :cat:", "NinoHeaders/1.0")
		// the new header will be "NinoHeaders/1.0"
		c.Response().Header().Set("Brought to you by", "{Dm;}")
		return next(c)
	}
}
