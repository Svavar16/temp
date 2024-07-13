package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func init () {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/api/auth/callback?provider=google"),
	)
}

func main() {
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Works!!!")
	})
	
	e.GET("/api/auth/callback", func(c echo.Context) error {
		user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, user)
	})

	e.GET("/api/logout", func(c echo.Context) error {
		gothic.Logout(c.Response(), c.Request().Response.Request)
		c.Response().Header().Set("Location", "/")
		return c.NoContent(http.StatusOK)
	})


	e.GET("/api/auth", func(c echo.Context) error {
		if gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request()); err == nil {
			return c.JSON(http.StatusOK, gothUser)
		} else {
			gothic.BeginAuthHandler(c.Response(), c.Request())
		}
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
