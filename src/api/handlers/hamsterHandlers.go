package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// struct for the POST handler for `hamsters`
type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// POST handler for `hamsters`
func AddHamster(h echo.Context) error { // easier to write but a bit slower
	hamster := Hamster{}
	err := h.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing 'addHamster' request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}
	log.Printf("This is your chubby hamster: %#v", hamster)
	return h.String(http.StatusOK, "We got your little hamster :)")
}
