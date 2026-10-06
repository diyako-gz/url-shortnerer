package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongo(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	mongoConncetionCheck := client.Ping(ctx, readpref.Primary())
	if mongoConncetionCheck != nil {
		return nil, mongoConncetionCheck
	}

	return client, err

}

func DisconnectMongo(client *mongo.Client) error {
	return client.Disconnect(context.Background())
}

func GetCollection(client *mongo.Client, dbName, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
