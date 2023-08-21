package db

import (
	"api_book/config"
	"api_book/helpers"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	Client *mongo.Client
}

func (m *MongoDatabase) RunConnect() *mongo.Client {
	uri := config.Get("MONGO_URL")
	if uri == "" {
		helpers.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		helpers.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		helpers.Fatal("Connection Failed to Mongo Database")
		helpers.Fatal(err)
	}
	return client
}

func (m *MongoDatabase) Setup(c *mongo.Client) {
	m.Client = c
}
