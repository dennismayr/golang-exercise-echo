package main

import (
	"fmt"

	"blueBot_go_webserver_echo/src/router"
)

// Main program
func main() {
	fmt.Println("Welcome to this humble server")

	// Server start
	echoInstance := router.New()
	echoInstance.Start(":8000")
}
