package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Hooray, you are on the secret admin/main page!")
}
