package repositories

import (
	"context"
	"fmt"
	"time"

	"github/Chidi-creator/go-medic-server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HospitalRepository interface {
	CreateHospital(ctx context.Context, hospital *models.Hospital) (*models.Hospital, error)
	GetHospitalById(ctx context.Context, id string) (*models.Hospital, error)
	GetAllHospitals(ctx context.Context) ([]models.Hospital, error)
	GetHospitalsByQuery(ctx context.Context, filter bson.M) ([]models.Hospital, error)
	UpdateHospitalById(ctx context.Context, id string, updateQuery bson.M) (*models.Hospital, error)
	DeleteHospital(ctx context.Context, id string) (int64, error)
}

// hospitalRepostory implements HospitalRepository
type hospitalRepository struct {
	client         *mongo.Client
	dbName         string
	collectionName string
}

// function that returns new Hospital Repository with necessary arguments
func NewHospitalRepository(client *mongo.Client, dbName string, collectionName string) HospitalRepository {

	collection := client.Database(dbName).Collection(collectionName)

	//Ensuring location functions & filters work

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"location.point": "2dsphere"},
		Options: options.Index().SetName("location_point_idx"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)

	if err != nil {
		fmt.Printf("Failed to create geo index for hospitals: %v\n", err)
	}

	return &hospitalRepository{
		client:         client,
		dbName:         dbName,
		collectionName: collectionName,
	}
}

//function that creates one hospital

func (h *hospitalRepository) CreateHospital(ctx context.Context, hospital *models.Hospital) (*models.Hospital, error) {
	collection := h.client.Database(h.dbName).Collection(h.collectionName)
	hospital.CreatedAt = time.Now()
	res, err := collection.InsertOne(ctx, hospital)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %w", err)
	}

	hospital.ID = res.InsertedID.(primitive.ObjectID)

	return hospital, nil
}

// function that retrieves hospital by Id
func (h *hospitalRepository) GetHospitalById(ctx context.Context, id string) (*models.Hospital, error) {
	collection := h.client.Database(h.dbName).Collection(h.collectionName)
	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": _id}

	var hospital models.Hospital
	err := collection.FindOne(ctx, filter).Decode(&hospital)

	if err != nil {
		return nil, fmt.Errorf("failed to find hospital: %w", err)
	}

	return &hospital, nil
}

// function that gets all hospital
func (h *hospitalRepository) GetAllHospitals(ctx context.Context) ([]models.Hospital, error) {
	collection := h.client.Database(h.dbName).Collection(h.collectionName)

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, fmt.Errorf("failed to find hospitals: %w", err)
	}
	defer cur.Close(ctx)
	var hospitals []models.Hospital

	for cur.Next(ctx) {
		var hospital models.Hospital
		err := cur.Decode(&hospital)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %w", err)
		}
		hospitals = append(hospitals, hospital)
	}
	return hospitals, nil

}

// function that gets hospitals by flexible query
func (h *hospitalRepository) GetHospitalsByQuery(ctx context.Context, filter bson.M) ([]models.Hospital, error) {
	collection := h.client.Database(h.dbName).Collection(h.collectionName)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find hospitals: %w", err)
	}
	defer cur.Close(ctx)

	var hospitals []models.Hospital
	for cur.Next(ctx) {
		var hospital models.Hospital
		err := cur.Decode(&hospital)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %w", err)
		}
		hospitals = append(hospitals, hospital)

	}
	return hospitals, nil

}

// function that updates hospitals by id
func (h *hospitalRepository) UpdateHospitalById(ctx context.Context, id string, updateQuery bson.M) (*models.Hospital, error) {
	collection := h.client.Database(h.dbName).Collection(h.collectionName)
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateQuery}

	var updatedResult models.Hospital
	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedResult)

	if err != nil {
		return nil, fmt.Errorf("failed to find hospitals: %w", err)
	}

	return &updatedResult, nil

}

//function that deletes hospital

func (h *hospitalRepository) DeleteHospital(ctx context.Context, id string) (int64, error) {
	collection := h.client.Database(h.dbName).Collection(h.collectionName)
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to delete hospital hospitals: %w", err)
	}
	return res.DeletedCount, nil

}
