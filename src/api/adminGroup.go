package api

import (
	"blueBot_go_webserver_echo/src/api/handlers"

	"github.com/labstack/echo/v4"
)

// `main` available under `/admin/main`
func AdminGroup(adminGroup *echo.Group) {
	adminGroup.GET("/main", handlers.MainAdmin)
}
