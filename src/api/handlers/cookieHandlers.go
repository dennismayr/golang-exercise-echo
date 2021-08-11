package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "You are now on the secret cookie page!")
}
