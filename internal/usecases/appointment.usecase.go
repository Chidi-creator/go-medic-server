package usecases

import (
	"context"
	"github/Chidi-creator/go-medic-server/internal/models"
	"github/Chidi-creator/go-medic-server/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type AppointmentUsecase interface {
	CreateAppointment(ctx context.Context, details *models.Appointment) (*models.Appointment, error)
	GetSingleAppointmentById(ctx context.Context, id string) (*models.Appointment, error)
	GetAppointmentsByDoctorId(ctx context.Context, id string) ([]models.Appointment, error)
	GetAppointmentsByUserId(ctx context.Context, id string) ([]models.Appointment, error)
	GetAppointmentsByQuery(ctx context.Context, filter bson.M) ([]models.Appointment, error)
	UpdateAppointmentById(ctx context.Context, id string, updateQuery bson.M) (*models.Appointment, error)
	DeleteAppointmentById(ctx context.Context, id string) (int64, error)
}

type appointmentUsecase struct{
	appointmentRepo repositories.AppointmentRepository
}

func NewAppointmentUsecase(appointmentRepo repositories.AppointmentRepository) AppointmentUsecase {
	return &appointmentUsecase{
		appointmentRepo: appointmentRepo,
	}
}

func (a *appointmentUsecase) CreateAppointment(ctx context.Context, details *models.Appointment) (*models.Appointment, error) {
	return a.appointmentRepo.CreateAppointment(ctx, details)
}

func (a *appointmentUsecase) GetSingleAppointmentById(ctx context.Context, id string) (*models.Appointment, error) {
	return a.appointmentRepo.GetSingleAppointmentById(ctx, id)
}

func (a *appointmentUsecase) GetAppointmentsByDoctorId(ctx context.Context, id string) ([]models.Appointment, error) {
	return a.appointmentRepo.GetAppointmentsByDoctorId(ctx, id)
}
func (a *appointmentUsecase) GetAppointmentsByUserId(ctx context.Context, id string) ([]models.Appointment, error) {
	return a.appointmentRepo.GetAppointmentsByUserId(ctx, id)
}
func (a *appointmentUsecase) GetAppointmentsByQuery(ctx context.Context, filter bson.M) ([]models.Appointment, error) {
	return a.appointmentRepo.GetAppointmentsByQuery(ctx, filter)
}
func (a *appointmentUsecase) UpdateAppointmentById(ctx context.Context, id string, updateQuery bson.M) (*models.Appointment, error) {
	return a.appointmentRepo.UpdateAppointmentById(ctx, id, updateQuery)
}
func (a *appointmentUsecase) DeleteAppointmentById(ctx context.Context, id string) (int64, error) {
	return a.appointmentRepo.DeleteAppointmentById(ctx, id)
}