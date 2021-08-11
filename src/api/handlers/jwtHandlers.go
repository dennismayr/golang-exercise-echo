package handlers

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Handler for mainJwt
func MainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	log.Println("User name: ", claims["name"], "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "You are at the Top Secret JWT section :)")
}
