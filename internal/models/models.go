package models

import (
	"github/Chidi-creator/go-medic-server/internal/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hospital struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Location  *Location          `json:"location,omitempty" bson:"location,omitempty"`
	Open      bool               `json:"open,omitempty" bson:"open,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Location struct {
	Address string    `json:"address" bson:"address"`
	Point   *GeoPoint `json:"point" bson:"point"`
}

type GeoPoint struct {
	Type        string    `bson:"type" json:"type"`               // always "Point"
	Coordinates []float64 `bson:"coordinates" json:"coordinates"` // [lon, lat]
}

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"passwprd,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Doctor struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Specialty  utils.Specialty    `json:"specialty,omitempty" bson:"specialty,omitempty"`
	HospitalID primitive.ObjectID `json:"hospitalId,omitempty" bson:"hospitalId,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Appointment struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	HospitalID primitive.ObjectID `json:"hospitalId,omitempty" bson:"hospitalId,omitempty"`
	UserID     primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	DoctorID   primitive.ObjectID `json:"doctorId,omitempty" bson:"doctorId,omitempty"`
	Status     utils.Status       `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
