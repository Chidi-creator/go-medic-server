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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppointmentRepository interface {
	CreateAppointment(ctx context.Context, details *models.Appointment) (*models.Appointment, error)
	GetSingleAppointmentById(ctx context.Context, id string) (*models.Appointment, error)
	GetAppointmentsByDoctorId(ctx context.Context, id string) ([]models.Appointment, error)
	GetAppointmentsByUserId(ctx context.Context, id string) ([]models.Appointment, error)
	GetAppointmentsByQuery(ctx context.Context, filter bson.M) ([]models.Appointment, error)
	UpdateAppointmentById(ctx context.Context, id string, updateQuery bson.M) (*models.Appointment, error)
	DeleteAppointmentById(ctx context.Context, id string) (int64, error)
}
type appointmentRepository struct {
	client     *mongo.Client
	dbName     string
	collection string
}

func NewAppointmentRepository(client *mongo.Client, dbName string, collection string) AppointmentRepository {
	return &appointmentRepository{
		client:     client,
		dbName:     dbName,
		collection: collection,
	}
}

func (a *appointmentRepository) CreateAppointment(ctx context.Context, details *models.Appointment) (*models.Appointment, error) {

	details.CreatedAt = time.Now()
	details.UpdatedAt = time.Now()

	collection := a.client.Database(a.dbName).Collection(a.collection)

	appointment, err := collection.InsertOne(ctx, details)
	if err != nil {
		return nil, fmt.Errorf("error creating appointment: %w", err)
	}
	details.ID = appointment.InsertedID.(primitive.ObjectID)

	return details, nil
}

func (a *appointmentRepository) GetSingleAppointmentById(ctx context.Context, id string) (*models.Appointment, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	filter := bson.M{"_id": _id}
	collection := a.client.Database(a.dbName).Collection(a.collection)

	var appointment models.Appointment
	err = collection.FindOne(ctx, filter).Decode(&appointment)

	if err != nil {
		return nil, fmt.Errorf("could not find appointment")
	}
	return &appointment, nil

}

func (a *appointmentRepository) GetAppointmentsByDoctorId(ctx context.Context, id string) ([]models.Appointment, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	filter := bson.M{"_id": _id}
	collection := a.client.Database(a.dbName).Collection(a.collection)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not find appointment with ID %v: %w", id, err)
	}
	defer cur.Close(ctx)
	var appointments []models.Appointment

	for cur.Next(ctx) {
		var appointment models.Appointment
		err := cur.Decode(&appointment)
		if err != nil {
			return nil, fmt.Errorf("cursor error")
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (a *appointmentRepository) GetAppointmentsByUserId(ctx context.Context, id string) ([]models.Appointment, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	filter := bson.M{"_id": _id}
	collection := a.client.Database(a.dbName).Collection(a.collection)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not find appointment with ID %v: %w", id, err)
	}
	defer cur.Close(ctx)

	var appointments []models.Appointment
	for cur.Next(ctx) {
		var appointment models.Appointment
		err := cur.Decode(&appointment)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %w", err)
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

func (a *appointmentRepository) FindAppointments(ctx context.Context, filter bson.M) ([]models.Appointment, error) {
	collection := a.client.Database(a.dbName).Collection(a.collection)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not find appointments: %w", err)
	}
	defer cur.Close(ctx)

	var appointments []models.Appointment
	for cur.Next(ctx) {
		var appointment models.Appointment
		err := cur.Decode(&appointment)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %w", err)
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

func (a *appointmentRepository) UpdateAppointmentById(ctx context.Context, id string, updateQuery bson.M) (*models.Appointment, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	filter := bson.M{"_id": _id}
	updateQuery["updatedAt"] = time.Now()

	update := bson.M{"$set": updateQuery}

	var updatedResult models.Appointment
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	collection := a.client.Database(a.dbName).Collection(a.collection)

	err = collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedResult)
	if err != nil {
		return nil, fmt.Errorf("failed to find hospitals: %w", err)
	}

	return &updatedResult, nil

}

func (a *appointmentRepository) DeleteAppointmentById(ctx context.Context, id string) (int64, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("invalid id: %w", err)
	}
	filter := bson.M{"_id": _id}

	collection := a.client.Database(a.dbName).Collection(a.collection)

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("could not delete record with %v: %w", id, err)
	}

	if res.DeletedCount == 0 {
		log.Println("No record was deleted")
	}

	return res.DeletedCount, err
}

func (a *appointmentRepository) GetAppointmentsByQuery(ctx context.Context, filter bson.M) ([]models.Appointment, error) {
	collection := a.client.Database(a.dbName).Collection(a.collection)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not find appointments: %w", err)
	}
	defer cur.Close(ctx)

	var appointments []models.Appointment
	for cur.Next(ctx) {
		var appointment models.Appointment
		err := cur.Decode(&appointment)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %w", err)
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}


