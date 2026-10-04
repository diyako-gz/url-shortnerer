package server

import (
	"net/http"
	"urlShortnerer/cmd/api/router"
	"urlShortnerer/cmd/config"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Start() error {
	ServerAddress := config.LoadServerAdd()

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "server run"})
	})

	v1 := e.Group("/api/v1")
	{
		health := v1.Group("/health")
		router.HealthRoute(health)
	}

	err := e.Start(ServerAddress)

	return err
}
