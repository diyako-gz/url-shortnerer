package main

import (
	"fmt"
	"log"
	"urlShortnerer/cmd/config"
	"urlShortnerer/cmd/db"
	"urlShortnerer/cmd/server"
)

func main() {
	config.LoadedEnv()

	mongoConfig := config.LoadMongo()
	mongoClient, err := db.ConnectMongo(mongoConfig.Url)
	if err != nil {
		log.Fatal("failed to connect to MongoDB:", err)
	} else {
		fmt.Println("connected to mongo")
	}

	mongo := mongoClient.Database(mongoConfig.DbName)
	log.Println("mongo connected", mongo.Name())

	defer func() {
		if err := db.DisconnectMongo(mongoClient); err != nil {
			log.Fatal("failed to disconnect from MongoDB:", err)
		}
		fmt.Println("disconnected from mongo")
	}()

	if err := server.Start(); err != nil {
		fmt.Println("failed to start server", "error:", err)
	}

}
