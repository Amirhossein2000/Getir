package storage

import (
	"Getir/internal/ports"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type repo struct {
	collection *mongo.Collection
}

func NewRepo(connection *mongo.Collection) ports.Storage {
	return &repo{
		connection,
	}
}

func (r *repo) SelectByTimeAndCount(startDate, endDate time.Time) ([]*ports.Record, error) {
	filter := bson.M{
		"createdAt": bson.M{
			"$gte": startDate.String(),
			"$lte": endDate.String(),
		},
	}

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	records := []*ports.Record{}
	for cursor.Next(context.Background()) {
		r := &ports.Record{}
		err := cursor.Decode(r)
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}

	return records, nil
}
