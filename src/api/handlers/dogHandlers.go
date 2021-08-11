package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// struct for the POST handler for `dogs`
type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// POST handler for `dogs`
func AddDog(d echo.Context) error { // this is the second-fastest method
	dog := Dog{}
	defer d.Request().Body.Close()
	err := json.NewDecoder(d.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing `addDog` request: %s", err)
		// return d.String(http.StatusInternalServerError, "")
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}
	log.Printf("This is your dog: %#v", dog)
	return d.String(http.StatusOK, "We got your doggie!")
}
