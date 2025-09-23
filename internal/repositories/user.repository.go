package repositories

import (
	"context"
	"fmt"
	"github/Chidi-creator/go-medic-server/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUsersByQuery(ctx context.Context, filter bson.M) ([]models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	UpdateUserById(ctx context.Context, id string, updateQuery bson.M) error
	DeleteUserById(ctx context.Context, id string) (int64, error)
}

type userRepository struct {
	client         *mongo.Client
	dbName         string
	collectionName string
}

func NewUserRepository(client *mongo.Client, dbName string, collectionName string) UserRepository {
	return &userRepository{
		client:         client,
		dbName:         dbName,
		collectionName: collectionName,
	}
}

func (u *userRepository) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	collection := u.client.Database(u.dbName).Collection(u.collectionName)

	user.CreatedAt = time.Now()

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not insert user: %w", err)
	}
	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil

}

func (u *userRepository) GetUsersByQuery(ctx context.Context, filter bson.M) ([]models.User, error) {
	collection := u.client.Database(u.dbName).Collection(u.collectionName)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not find users: %w", err)
	}
	defer cur.Close(ctx)
	var users []models.User
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %w", err)
		}
		users = append(users, user)
	}

	return users, nil

}

func (u *userRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}
	collection := u.client.Database(u.dbName).Collection(u.collectionName)

	filter := bson.M{"_id": _id}

	var user models.User

	err = collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, fmt.Errorf("could not find User: %w", err)
	}

	return &user, nil

}

func (u *userRepository) UpdateUserById(ctx context.Context, id string, updateQuery bson.M) error {
	collection := u.client.Database(u.dbName).Collection(u.collectionName)

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	updateQuery["updatedAt"] = time.Now()

	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateQuery}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no documents matched: %w", err)
	}

	return nil

}

func (u *userRepository) DeleteUserById(ctx context.Context, id string) (int64, error) {
	collection := u.client.Database(u.dbName).Collection(u.collectionName)

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID: %w", err)
	}

	filter := bson.M{"_id": _id}

	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return 0, fmt.Errorf("could not delete user : %w", err)
	}
	return res.DeletedCount, nil
}
