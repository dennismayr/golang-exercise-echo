package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// struct for the POST handler for `cats`
type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// POST handler for `cats`
func AddCat(c echo.Context) error { // this method is the fastest (barebones, `ioutil`-based)
	cat := Cat{}
	defer c.Request().Body.Close()             // to handle the Body we need to close this first
	b, err := ioutil.ReadAll(c.Request().Body) // b=body
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &cat) // the body value `b` will be stored in the `&cat` pointer -> memory address of `cat` (Cat{})
	if err != nil {
		log.Printf("Failed parsing body for 'addCats': %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("This is your cat: %#v", cat)
	return c.String(http.StatusOK, "We got your cat!")
}

// Exercise: cats
func GetCats(c echo.Context) error {
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
