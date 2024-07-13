package web

import (
	"embed"

	"github.com/labstack/echo/v4"
)

type Web struct{}

var (
	//go:embed all:dist
	dist embed.FS
	//go:embed dist/index.html
	indexHTML embed.FS

	distDirFS     = echo.MustSubFS(dist, "dist")
	distIndexHtml = echo.MustSubFS(indexHTML, "dist")

	routes = []string{"/", "/bla", "/bla/:slug"}
)

func (rh *Web) RegisterHandlers(e *echo.Echo) {
	for _, route := range routes {
		e.FileFS(route, "index.html", distIndexHtml)
	}
	e.StaticFS("/", distDirFS)
}
