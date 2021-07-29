package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, folks")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	return c.String(http.StatusOK, fmt.Sprintf("Your cat name is: %s \nand his type is: %s\n", catName, catType))
}

//Main program
func main() {
	fmt.Println("Welcome to this humble server")
	e := echo.New()

	// Routes
	e.GET("/", hello)
	e.GET("/cats", getCats)

	//Server start
	e.Start(":8000")
}
