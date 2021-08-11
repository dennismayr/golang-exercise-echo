package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAdminMiddleware(adminGroup *echo.Group) {
	// Logging the server's interactions
	adminGroup.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: `[${time_rfc3339}] | (${protocol}) ${status} ${method} ${host}${path} | ${latency_human}` + "\n",
		})) // Format: `literal order, you can do whatever you want`

	// Adding Basic Authentication to the `admin` group
	adminGroup.Use(middleware.BasicAuth(
		func(username, password string, c echo.Context) (bool, error) {
			// Check DB if pw is valid
			if username == "dmayr" && password == "1234" {
				return true, nil
			}
			return true, nil
		}))
}
