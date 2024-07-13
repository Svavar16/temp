package main

import (
	web "reactecho/client"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	web.RegisterHandlers(e)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Works!!")
	// })

	e.Logger.Fatal(e.Start(":8080"))
}
