package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func SetCookieMiddleware(g *echo.Group) {
	g.Use(checkCookie)
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
