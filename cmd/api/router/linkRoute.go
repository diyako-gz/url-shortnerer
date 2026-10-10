package router

import (
	"urlShortnerer/cmd/api/handler"

	"github.com/labstack/echo/v5"
)

func LinkRoute(g *echo.Group) {
	g.POST("", handler.CreateLinkHandler)
	g.GET("/:id", handler.GetLinkHandler)
}
