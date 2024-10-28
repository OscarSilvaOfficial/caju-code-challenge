package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB[Data any] struct {
	client     *mongo.Client
	database   string
}

func (db *MongoDB[Data]) Find(collectionOrTable string, where map[string]interface{}) ([]Data, error) {
	var results []Data
	collection := db.client.Database(db.database).Collection(collectionOrTable)

	filter := bson.M(where)
	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		return nil, fmt.Errorf("failed to execute Find: %v", err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result Data

		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode result: %v", err)
		}

		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return results, nil
}

func (db *MongoDB[Data]) Insert(collectionOrTable string, data Data) (interface{}, error) {
	collection := db.client.Database(db.database).Collection(collectionOrTable)
	result, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)
	}
	return result.InsertedID, nil
}

func NewMongoDB[Data any](uri string, dbName string) (*MongoDB[Data], error) {
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	return &MongoDB[Data]{client: client, database: dbName}, nil
}
