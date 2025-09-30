package repositories

import (
	"context"
	"fmt"
	"github/Chidi-creator/go-medic-server/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DoctorRepository interface {
	CreateDoctor(ctx context.Context, doctor *models.Doctor) (*models.Doctor, error)
	FindDoctorById(ctx context.Context, id string) (*models.Doctor, error)
	FindDoctorsByQuery(ctx context.Context, filter bson.M) ([]models.Doctor, error)
	GetDoctorsByHospitalId(ctx context.Context, id string) ([]models.Doctor, error)
	UpdateDoctorById(ctx context.Context, id string, updateQuery bson.M) error
	DeleteDoctorByUserId(ctx context.Context, id string) error
}

type doctorRepository struct {
	Client     *mongo.Client
	dbName     string
	collection string
}

func NewDoctorRepository(client *mongo.Client, dbName string, collection string) DoctorRepository {
	return &doctorRepository{
		Client:     client,
		dbName:     dbName,
		collection: collection,
	}
}

func (d *doctorRepository) CreateDoctor(ctx context.Context, doctor *models.Doctor) (*models.Doctor, error) {
	collection := d.Client.Database(d.dbName).Collection(d.collection)

	doctor.CreatedAt = time.Now()
	doctor.UpdatedAt = time.Now()

	res, err := collection.InsertOne(ctx, doctor)
	if err != nil {
		return nil, fmt.Errorf("could not insert record: %w", err)
	}
	fmt.Println("inserted record with id of: ", res.InsertedID)

	doctor.ID = res.InsertedID.(primitive.ObjectID)

	return doctor, nil
}

func (d *doctorRepository) FindDoctorById(ctx context.Context, id string) (*models.Doctor, error) {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid Id: %w", err)
	}
	collection := d.Client.Database(d.dbName).Collection(d.collection)

	filter := bson.M{"_id": _id}

	var doctor models.Doctor
	err = collection.FindOne(ctx, filter).Decode(&doctor)

	if err != nil {
		return nil, fmt.Errorf("could not find user: %w", err)
	}

	return &doctor, nil

}

func (d *doctorRepository) GetDoctorsByHospitalId(ctx context.Context, id string) ([]models.Doctor, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid object id: %w", err)
	}
	filter := bson.M{"hospitalId": _id}

	collection := d.Client.Database(d.dbName).Collection(d.collection)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not find doctors: %w", err)
	}
	defer cur.Close(ctx)
	var doctors []models.Doctor
	for cur.Next(ctx) {
		var doctor models.Doctor
		err := cur.Decode(&doctor)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %w", err)
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}
func (d *doctorRepository) FindDoctorsByQuery(ctx context.Context, filter bson.M) ([]models.Doctor, error) {
	collection := d.Client.Database(d.dbName).Collection(d.collection)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not find doctors: %w", err)
	}
	defer cur.Close(ctx)
	var doctors []models.Doctor
	for cur.Next(ctx) {
		var doctor models.Doctor
		err := cur.Decode(&doctor)
		if err != nil {
			return nil, fmt.Errorf("cursor err")
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil

}

func (d *doctorRepository) UpdateDoctorById(ctx context.Context, id string, updateQuery bson.M) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid object id: %w", err)
	}
	collection := d.Client.Database(d.dbName).Collection(d.collection)

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

	log.Printf("A total number of %v was updated", res.ModifiedCount)
	return nil

}

func (d *doctorRepository) DeleteDoctorByUserId(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid object id: %w", err)
	}
	filter := bson.M{"_id": _id}

	collection := d.Client.Database(d.dbName).Collection(d.collection)
	res, err := collection.DeleteOne(ctx, filter)

	if res.DeletedCount == 0 {
		log.Println("No record was deleted")
	}
	if err != nil {
		return fmt.Errorf("could not delete doctor by id")
	}
	return nil

}
