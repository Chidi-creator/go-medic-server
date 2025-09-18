package mongo

import (
	"context"
	"fmt"
	"github/Chidi-creator/go-medic-server/config"
	"time"

	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)



func ConnectMongo() {

	connectionString := config.AppConfig.Mongo_URI

	// Validate connection string
	if connectionString == "" {
		log.Fatal("MongoDB connection string is empty")
	}

	fmt.Printf("Attempting to connect to MongoDB...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDb
	var err error
	client, err = mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	fmt.Printf("Pinging MongoDB server...\n")
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed (connection string might be invalid):", err)
	}
	hospitalCollection := client.Database(config.AppConfig.DB_NAME).Collection("hospitals")
	hospitalCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "location.point", Value: "2dsphere"}},
	})

	fmt.Println("MongoDB connection successful")

}

// //helper functions

func InsertOne[T any](ctx context.Context, collectionName string, document T) (*T, error) {

	collection := client.Database(config.AppConfig.DB_NAME).Collection(collectionName)

	_, err := collection.InsertOne(ctx, document)

	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %w", err)

	}

	return &document, nil

}

func FindMany[T any](ctx context.Context, collectionName string, filter bson.M) ([]T, error) {

	collection := client.Database(config.AppConfig.DB_NAME).Collection(collectionName)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find documents in %s: %w", collectionName, err)
	}
	defer cur.Close(ctx)

	var results []T

	for cur.Next(ctx) {
		var result T
		err := cur.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("failed to decode document: %w", err)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return results, nil

}

func FindOneById[T any](ctx context.Context, collectionName string, id string) (*T, error) {
	collection := client.Database(config.AppConfig.DB_NAME).Collection(collectionName)

	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": _id}
	var result T
	err := collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return nil, fmt.Errorf("failed to decode document: %w", err)
	}

	return &result, nil

}

func FindOne[T any](ctx context.Context, collectionName string, filter bson.M) (*T, error) {
	collection := client.Database(config.AppConfig.DB_NAME).Collection(collectionName)

	res := collection.FindOne(ctx, filter)
	var result T
	if err := res.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode document: %w", err)
	}

	return &result, nil

}

func DeleteOne(ctx context.Context, collectionName string, filter bson.M) error {
	collection := client.Database(config.AppConfig.DB_NAME).Collection(collectionName)

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete document from %s: %w", collectionName, err)
	}

	return nil
}

func UpdateOne[T any](ctx context.Context, collectionName string, filter bson.M, updateQuery bson.M) (*T, error) {
	collection := client.Database(config.AppConfig.DB_NAME).Collection(collectionName)

	result := collection.FindOneAndUpdate(ctx, filter, updateQuery)

	var updated T
	if err := result.Decode(&updated); err != nil {
		return nil, fmt.Errorf("failed to decode updated document: %w", err)
	}

	return &updated, nil

}
