package api

import (
	"blueBot_go_webserver_echo/src/api/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(echoInstance *echo.Echo) {
	// Routes
	echoInstance.GET("/hello", handlers.Hello)
	echoInstance.GET("/cats/:data", handlers.GetCats)
	echoInstance.GET("/login", handlers.Login)

	// Endpoint for posting
	echoInstance.POST("/cats", handlers.AddCat)
	echoInstance.POST("/dogs", handlers.AddDog)
	echoInstance.POST("/hamsters", handlers.AddHamster)
}
