package config

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    mongoClient *mongo.Client
)

func DbConfig() *mongo.Client {
	// Connect to MongoDB
	println("")
	clientOptions := options.Client().ApplyURI(GetEnv("MONGO_CONNECTION_STRING"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	mongoClient = client
	return nil
}

func GetMongoClient() *mongo.Client {
    return mongoClient
}
