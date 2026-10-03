package server

import (
	"net/http"
	"urlShortnerer/src/api/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Start() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "server run"})
	})

	serverAddr, err := godotenv.Read("../.env")

	if err != nil {
		e.Logger.Error("cannot read env file", err)
	}

	v1 := e.Group("/api/v1")
	{
		health := v1.Group("/health")
		router.HealthRoute(health)
	}

	if err = e.Start(serverAddr["ServerAddress"]); err != nil {
		e.Logger.Error("faild to start server", err)
	}

}
