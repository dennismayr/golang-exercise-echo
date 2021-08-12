package main

import (
	"blueBot_go_webserver_echo/src/router"
	"fmt"
)

// Main program
func main() {
	fmt.Println("Welcome to this humble server")

	// Server start
	echoInstance := router.New()
	echoInstance.Start(":8000")
}
