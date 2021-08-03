package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// Echo framework
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Test function
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from this new Echo server! Welcome!")
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

// struct for the POST handler for `cats`
type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// struct for the POST handler for `dogs`
type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// struct for the POST handler for `hamsters`
type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// POST handler for `cats`
func addCat(c echo.Context) error { // this method is the fastest (barebones, `ioutil`-based)
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

// POST handler for `dogs`
func addDog(d echo.Context) error { // this is the second-fastest method
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

// POST handler for `hamsters`
func addHamster(h echo.Context) error { // easier to write but a bit slower
	hamster := Hamster{}
	err := h.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing 'addHamster' request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}
	log.Printf("This is your chubby hamster: %#v", hamster)
	return h.String(http.StatusOK, "We got your little hamster :)")
}

// Handler for `mainAdmin`
func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Hooray, you are on the secret admin/main page!")
}

// Main program
func main() {
	fmt.Println("Welcome to this humble server")
	e := echo.New()

	// Grouping: Middleware exercise
	g := e.Group("/admin") // group root

	// Logging the server's interactions
	g.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: `[${time_rfc3339}] | (${protocol}) ${status} ${method} ${host}${path} | ${latency_human}` + "\n",
		})) // Format: `literal order, you can do whatever you want`

	// Adding Basic Authentication to the `admin` group
	g.Use(middleware.BasicAuth(
		func(username, password string, c echo.Context) (bool, error) {
			// Check DB if pw is valid
			if username == "dmayr" && password == "1234" {
				return true, nil
			}
			return false, nil
		}))

	// Routes
	e.GET("/", hello)
	e.GET("/cats/:data", getCats)
	// `main` available under `/admin/main`
	g.GET("/main", mainAdmin)

	// Endpoint for posting
	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/hamsters", addHamster)

	// Server start
	e.Start(":8000")
}
