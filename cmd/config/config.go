package config

import (
	"os"

	"github.com/joho/godotenv"
)

type MongoDb struct {
	DbName string
	Url    string
}

type RedisDb struct {
	Addr string
}

type LoadEnv struct {
	Env string
}

var ServerAddress string

func LoadedEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		print("error ignored")
	}
}

func LoadServerAdd() (serverAddres string) {

	ServerAddress = os.Getenv("ServerAddress")

	if ServerAddress == "" {
		ServerAddress = ":5050"
		return ServerAddress
	}

	return ServerAddress

}

func LoadMongo() *MongoDb {
	return &MongoDb{
		DbName: os.Getenv("MongoName"),
		Url:    os.Getenv("MongoUrl"),
	}
}

func LoadRedis() *RedisDb {
	return &RedisDb{
		Addr: os.Getenv("RedisAddr"),
	}
}
