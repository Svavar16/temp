package webAdmin

import (
	"embed"

	"github.com/labstack/echo/v4"
)

type AdminWeb struct{}

var (
	//go:embed all:dist
	dist embed.FS
	//go:embed dist/index.html
	indexHTML embed.FS

	distDirFS     = echo.MustSubFS(dist, "dist")
	distIndexHtml = echo.MustSubFS(indexHTML, "dist")

	routes = []string{"/"}
)

func (rah *AdminWeb) RegisterAdminHandlers(e *echo.Echo) {
	for _, route := range routes {
		e.FileFS(route, "index.html", distIndexHtml)
	}
	e.StaticFS("/", distDirFS)
}
