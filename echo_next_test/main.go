package main

import (
	"log"
	web "nextjstest/client2"

	"github.com/labstack/echo/v4"
)



func main() {
	e := echo.New()

	web.RegisterHandlers(e)
	
	log.Println("works!")
	e.Logger.Fatal(e.Start(":8080"))
}
