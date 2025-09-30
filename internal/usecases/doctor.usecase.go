package usecases

import (
	"context"
	"github/Chidi-creator/go-medic-server/internal/models"
	"github/Chidi-creator/go-medic-server/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type DoctorUsecase interface {
	CreateDoctor(ctx context.Context, doctor *models.Doctor) (*models.Doctor, error)
	FindDoctorById(ctx context.Context, id string) (*models.Doctor, error)
	FindDoctorsByQuery(ctx context.Context, filter bson.M) ([]models.Doctor, error)
	GetDoctorsByHospitalId(ctx context.Context, id string) ([]models.Doctor, error)
	UpdateDoctorById(ctx context.Context, id string, updateQuery bson.M) error
	DeleteDoctorByUserId(ctx context.Context, id string) error
}

type doctorUsecase struct {
	doctorRepo repositories.DoctorRepository
}

func NewDoctorUseCase(doctorRepo repositories.DoctorRepository) DoctorUsecase {
	return &doctorUsecase{
		doctorRepo: doctorRepo,
	}
}

func (d *doctorUsecase) CreateDoctor(ctx context.Context, doctor *models.Doctor) (*models.Doctor, error) {
	newDoctor, err := d.doctorRepo.CreateDoctor(ctx, doctor)
	if err != nil {
		return nil, err
	}

	return newDoctor, nil
}
func (d *doctorUsecase) FindDoctorById(ctx context.Context, id string) (*models.Doctor, error) {
	doctor, err := d.doctorRepo.FindDoctorById(ctx, id)
	if err != nil {
		return nil, err
	}
	return doctor, nil
}
func (d *doctorUsecase) FindDoctorsByQuery(ctx context.Context, filter bson.M) ([]models.Doctor, error) {
	doctors, err := d.doctorRepo.FindDoctorsByQuery(ctx, filter)
	if err != nil {
		return nil, err
	}
	return doctors, nil
}
func (d *doctorUsecase) GetDoctorsByHospitalId(ctx context.Context, id string) ([]models.Doctor, error) {
	doctors, err := d.doctorRepo.GetDoctorsByHospitalId(ctx, id)
	if err != nil {
		return nil, err
	}
	return doctors, nil
}
func (d *doctorUsecase) UpdateDoctorById(ctx context.Context, id string, updateQuery bson.M) error {
	err := d.doctorRepo.UpdateDoctorById(ctx, id, updateQuery)
	if err != nil {
		return err
	}
	return nil
}
func (d *doctorUsecase) DeleteDoctorByUserId(ctx context.Context, id string) error {
	err := d.doctorRepo.DeleteDoctorByUserId(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
