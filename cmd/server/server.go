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
	"go.mongodb.org/mongo-driver/mongo"
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
		links := v1.Group("/links")
		router.LinkRoute(links)
	}

	err := e.Start(ServerAddress)

	return err
}

func Run() error {
	var MongoClient *mongo.Client
	config.LoadedEnv()

	mongoConfig := config.LoadMongo()
	MongoClient, err := db.ConnectMongo(mongoConfig.Url)
	if err != nil {
		fmt.Println("failed to connect to MongoDB:", err)
		return err
	} else {
		fmt.Println("connected to mongo")
	}

	mongo := MongoClient.Database(mongoConfig.DbName)
	log.Println("mongo connected", mongo.Name())

	defer func() {
		if err := db.DisconnectMongo(MongoClient); err != nil {
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
