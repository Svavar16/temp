package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	// this is something that could be used in future project, when I want to use subdomain
	// trying out the subdomain
	hosts := map[string]*Host{}

	// testing admin
	admin := echo.New()
	hosts["admin.localhost:8080"] = &Host{admin}

	admin.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "From Admin")
	})

	admin.GET("/:string", func(c echo.Context) error {
		result := c.Param("string")
		return c.String(http.StatusOK, "From Admin - "+result)
	})

	// testing normal
	site := echo.New()
	hosts["localhost:8080"] = &Host{site}

	site.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "From Normal")
	})

	site.GET("/:string", func(c echo.Context) error {
		result := c.Param("string")
		return c.String(http.StatusOK, "From Normal - "+result)
	})

	// the routing
	e := echo.New()
	e.GET("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

	e.Start(":8080")
}
