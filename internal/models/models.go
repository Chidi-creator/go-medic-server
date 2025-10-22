package models

import (
	"github/Chidi-creator/go-medic-server/internal/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hospital struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=2,max=100"`
	Location    *Location          `json:"location,omitempty" bson:"location,omitempty" validate:"required"`
	UserID      primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty" validate:"required"`
	Specialties []utils.Specialty  `json:"specialties,omitempty" bson:"specialties,omitempty" validate:"required,dive,required,specialties"`
	Open        bool               `json:"open,omitempty" bson:"open,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,e164"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Location struct {
	Address string    `json:"address" bson:"address" validate:"required"`
	Point   *GeoPoint `json:"point" bson:"point" validate:"required"`
}

type GeoPoint struct {
	Type        string    `bson:"type" json:"type" validate:"required,geopoint"`                    //always Point
	Coordinates []float64 `bson:"coordinates" json:"coordinates" validate:"required,dive,required"` //[longitude, latitude]
}

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required,min=2,max=100"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required,min=2,max=100"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=6"`
	Roles     []utils.Roles      `json:"roles,omitempty" bson:"roles,omitempty" validate:"required,min=1,dive,required,roles"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Doctor struct {
	ID           primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname    string              `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required,min=2,max=100"`
	LastName     string              `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required,min=2,max=100"`
	Specialties  []utils.Specialty   `json:"specialties,omitempty" bson:"specialties,omitempty" validate:"required,dive,required"`
	HospitalID   primitive.ObjectID  `json:"hospitalId,omitempty" bson:"hospitalId,omitempty" validate:"required"`
	UserID       *primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	InviteStatus utils.InviteStatus  `json:"inviteStatus,omitempty" bson:"inviteStatus,omitempty" validate:"oneof=pending accepted rejected"`
	CreatedAt    time.Time           `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    time.Time           `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Appointment struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	HospitalID primitive.ObjectID `json:"hospitalId,omitempty" bson:"hospitalId,omitempty"`
	UserID     primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	DoctorID   primitive.ObjectID `json:"doctorId,omitempty" bson:"doctorId,omitempty"`
	Status     utils.Status       `json:"status,omitempty" bson:"status,omitempty"`
	Reason     string             `json:"reason,omitempty" bson:"reason,omitempty" validate:"required,min=5,max=500"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
