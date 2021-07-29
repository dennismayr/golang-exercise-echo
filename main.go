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

func main() {
	fmt.Println("Welcome to this humble server")
	e := echo.New()
	e.GET("/", hello)

	e.Start(":8000")
}
