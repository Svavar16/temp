package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
