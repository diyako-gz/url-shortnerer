package server

import (
	"fmt"
	"log"
	"net/http"
	"urlShortnerer/cmd/api/router"
	"urlShortnerer/cmd/config"
	"urlShortnerer/cmd/db"

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

func Run() error {
	config.LoadedEnv()

	mongoConfig := config.LoadMongo()
	mongoClient, err := db.ConnectMongo(mongoConfig.Url)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	} else {
		fmt.Println("connected to mongo")
	}

	mongo := mongoClient.Database(mongoConfig.DbName)
	log.Println("mongo connected", mongo.Name())

	defer func() {
		if err := db.DisconnectMongo(mongoClient); err != nil {
			fmt.Println("failed to disconnect from MongoDB:", err)
		}
		fmt.Println("disconnected from mongo")
	}()

	redisConfig := config.LoadRedis()
	redisClient, err := db.ConnectRedis(redisConfig.Addr)
	if err != nil {
		return err
	} else {
		fmt.Println("connected to redis")
	}

	defer func() {
		if err := db.DisconnectRedis(redisClient); err != nil {
			fmt.Println("failed to disconnect from Redis:", err)
		}
		fmt.Println("disconnected from redis")
	}()

	if err := Start(); err != nil {
		return err
	}
	return err
}
