package router

import (
	"urlShortnerer/src/api/handler"

	"github.com/labstack/echo/v5"
)

func HealthRoute(e *echo.Group) {

	e.GET("/healthCheck", handler.HealthHandler)

}
