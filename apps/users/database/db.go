package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func InitDB(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to MongoDB!")
	return Client, nil
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("monorepo").Collection(collectionName)
}
