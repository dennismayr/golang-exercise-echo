package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// Handler for login
func Login(c echo.Context) error {
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

		// JWTCookie implementation
		jwtCookie := &http.Cookie{}
		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(jwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in successfully.",
			"token":   token,
		})
	}
	return c.String(http.StatusUnauthorized, "Your username or password don't match.")
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
