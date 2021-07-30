package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"net/http"
)

//Test function
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, folks")
}

// Exercise: cats
func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your cat name is: %s \nand his type is: %s\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "You need to let us know if you want JSON formatting or string data",
	})
}

//Main program
func main() {
	fmt.Println("Welcome to this humble server")
	e := echo.New()

	// Routes
	e.GET("/", hello)
	e.GET("/cats/:data", getCats)

	//Server start
	e.Start(":8000")
}
