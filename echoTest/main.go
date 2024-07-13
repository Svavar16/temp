package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var cookieSecret = "ThisIsMySecret"

func main() {
	e := echo.New()

	e.Use(changeCookieToAuth)

	e.GET("/", test)
	e.GET("/cookie", getCookie)
	e.GET("/cookie/read", readCookie)
	e.GET("/cookie/remove", removeCookie)
	e.GET("/private", private, isLoggedIn) // here we have the route that will be protected

	e.Logger.Fatal(e.Start(":8080"))
}

func test(e echo.Context) error{
	return e.String(http.StatusOK, "Hello, World!")
}

func getCookie(e echo.Context) error {
	claims := jwt.MapClaims{
		"name": "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(cookieSecret))
	if err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "test"
	cookie.Value = t
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(5 * time.Hour)
	e.SetCookie(cookie)
	return e.String(http.StatusOK, "Cookie Sent")
}

func readCookie(e echo.Context) error {
	cookie, err := e.Cookie("test")
	if err != nil {
		e.String(http.StatusUnauthorized, err.Error())
		// return err
	}
	log.Println(cookie)
	return e.String(http.StatusOK, "Cookie read")
}

func removeCookie(e echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "test"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.MaxAge = -1
	cookie.Expires = time.Now().Add(-100 * time.Hour)
	
	e.SetCookie(cookie)
	e.Request().Header.Del("Authorization")
	log.Println(e.Request().Header)

	return e.String(http.StatusOK, "Cookie removed")
}

// this is the function that should use the jwt
func private(c echo.Context) error {
	return c.String(http.StatusOK, "Private!!!")
}

func changeCookieToAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("test")
		if err != nil {
			c.Request().Header.Del("Authorization")
			return next(c)
		}
		bearerToken := fmt.Sprintf("Bearer %v", cookie.Value)
		log.Println(bearerToken)
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %v", cookie.Value))
		// c.Request().Header.Set("Authorization", bearerToken)
		log.Println(c.Request())

		return next(c)
	}
}

// example of how the jwt middleware could be
// this could be then plugged into the function
var isLoggedIn = echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte(cookieSecret), // should be a env variabled
	// TokenLookup: "cookie:test",
})