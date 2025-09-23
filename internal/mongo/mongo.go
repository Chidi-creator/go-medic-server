package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	Client *mongo.Client
}

//New client created a new MongoDB client wrapper

func NewClient(uri string, dbName string) (*Client, error) {

	if uri == "" {
		return nil, fmt.Errorf("MongoDB URI is required")

	}
	// Increase the context timeout for initial connection
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri).SetConnectTimeout(10 * time.Second).SetServerSelectionTimeout(15 * time.Second).SetMaxPoolSize(100).SetRetryWrites(true).SetRetryReads(true)

	//connect to mongo db
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %w", err)
	}

	// Create a separate context for ping to avoid timeout issues
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer pingCancel()

	// Ping the database to verify connection
	if err = client.Ping(pingCtx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return &Client{Client: client}, nil

}
