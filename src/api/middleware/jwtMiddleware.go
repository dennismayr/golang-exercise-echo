package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetJwtMiddleware(jwtGroup *echo.Group) {
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
		// Restriction scheme for Authorization
		TokenLookup: "cookie:JWTCookie", // instead of default "Authorization"
	}))
}
