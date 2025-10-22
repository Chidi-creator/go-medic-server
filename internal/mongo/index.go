package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes(client *mongo.Client, dbName string) {
	log.Println("Creating indexes...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := client.Database(dbName)

	//USERS INDEX
	userCollection := db.Collection("users")

	userIndex := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("unique_email_idx"),
		},
	}

	if _, err := userCollection.Indexes().CreateMany(ctx, userIndex); err != nil {
		fmt.Printf("Failed to create user index: %v", err)
	} else {
		fmt.Println("User index created successfully")
	}
}
