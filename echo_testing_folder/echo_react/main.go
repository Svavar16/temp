package main

import (
	"net/http"

	web "echoreact/client"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	web.RegisterHandlers(e)

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// C:\programming\test\echo_testing_folder\echo_react\client\build\index.html
