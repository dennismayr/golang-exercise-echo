package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
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

// struct for the JWT implementation
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
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

// Handler for `mainCookie`
func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "You've come to the secret cookie place :)")
}

// Handler for mainJwt
func mainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	log.Println("User name: ", claims["name"], "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "You are at the Top Secret JWT section :)")
}

// Handler for login
func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	// Check user and password against the DB, in the future
	if username == "dmayr" && password == "1234" {
		cookie := &http.Cookie{}
		// cookie := new(http.Cookie)
		cookie.Name = "sessionID"
		cookie.Value = "ninoCookie_id"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		// c.SetCookie(cookie *http.Cookie)
		c.SetCookie(cookie)

		// OK message as string:
		// return c.String(http.StatusOK, "You were logged in successfully.")
		// OK message in JSON:
		token, err := createJwtToken()
		if err != nil {
			log.Printf("Error when creating JWT token: %s", err)
			return c.String(http.StatusInternalServerError, "Something has gone wrong")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in successfully.",
			"token":   token,
		})
	}
	return c.String(http.StatusUnauthorized, "Your username or password don't match.")
}

//////////////////////////////////////////////////////////////////////
//                            Middlewares                           //
//////////////////////////////////////////////////////////////////////

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	// HandlerFunc defines a function to serve HTTP requests
	return func(c echo.Context) error {
		// c.Response().Header().Set(echo.HeaderServer, "NinoHeaders/1.0")
		c.Response().Header().Set("InterceptedBy", "NinoHeaders/1.0")
		// the new header will be "NinoHeaders/1.0"
		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			// Handler: print a better, readable, custom string instead of the bare error message, by matching and replacing it:
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "You don't seem to have any cookies, fam.")
			}
			log.Println(err)
			return err
		}
		if cookie.Value == "ninoCookie_id" {
			return next(c)
		}
		return c.String(http.StatusUnauthorized, "You don't seem to have the right cookie.")
	}
}

// JWT Token: external, yet accepted middleware
func createJwtToken() (string, error) {
	claims := JwtClaims{
		"dmayr",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecret")) // this should be a private key, never hardcode a string like here
	if err != nil {
		return "", err
	}
	return token, nil
}

//////////////////////////////////////////////////////////////////////

// Main program
func main() {
	fmt.Println("Welcome to this humble server")
	echoInstance := echo.New()

	// Adding the ServerHeader to the Group
	echoInstance.Use(ServerHeader)

	// GROUPS
	// Middleware exercise
	adminGroup := echoInstance.Group("/admin") // group root

	// Cookies
	cookieGroup := echoInstance.Group("/cookie")

	// JWT:
	jwtGroup := echoInstance.Group("/jwt")

	// Logging the server's interactions
	adminGroup.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: `[${time_rfc3339}] | (${protocol}) ${status} ${method} ${host}${path} | ${latency_human}` + "\n",
		})) // Format: `literal order, you can do whatever you want`

	// Adding Basic Authentication to the `admin` group
	adminGroup.Use(middleware.BasicAuth(
		func(username, password string, c echo.Context) (bool, error) {
			// Check DB if pw is valid
			if username == "dmayr" && password == "1234" {
				return true, nil
			}
			return false, nil
		}))

	// JWT enforcing
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
		// Restriction scheme for Authorization
		TokenLookup: "header:MyHeader", // instead of default "Authorization"
		AuthScheme:  "iLoveGatitos",    // Instead of default "Bearer"
	}))

	// HANDLERS
	cookieGroup.Use(checkCookie)
	// `main` available under `/admin/main`
	adminGroup.GET("/main", mainAdmin)
	// Cookie implementation
	cookieGroup.GET("/main", mainCookie)
	// JWT handler
	jwtGroup.GET("/main", mainJwt)

	// Routes
	echoInstance.GET("/", hello)
	echoInstance.GET("/cats/:data", getCats)
	echoInstance.GET("/login", login)

	// Endpoint for posting
	echoInstance.POST("/cats", addCat)
	echoInstance.POST("/dogs", addDog)
	echoInstance.POST("/hamsters", addHamster)

	// Server start
	echoInstance.Start(":8000")
}
