package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoParams struct {
	Address    string
	Database   string
	Collection string
}

func NewMongo(params *MongoParams) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(params.Address)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	connection := client.Database(params.Database).Collection(params.Collection)

	return connection, nil
}
