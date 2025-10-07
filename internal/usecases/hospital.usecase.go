package usecases

import (
	"context"
	"github/Chidi-creator/go-medic-server/internal/models"
	"github/Chidi-creator/go-medic-server/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type HospitalUsecase interface {
	CreateHospital(ctx context.Context, hospital *models.Hospital) (*models.Hospital, error)
	GetHospitalById(ctx context.Context, id string) (*models.Hospital, error)
	GetAllHospitals(ctx context.Context) ([]models.Hospital, error)
	GetHospitalsByQuery(ctx context.Context, filter bson.M) ([]models.Hospital, error)
	UpdateHospitalById(ctx context.Context, id string, updateQuery bson.M) (*models.Hospital, error)
	DeleteHospital(ctx context.Context, id string) (int64, error)
}

type hospitalUsecase struct {
	hospitalRepo repositories.HospitalRepository
}

func NewHospitalUseCase(hospitalRepo repositories.HospitalRepository) HospitalUsecase {
	return &hospitalUsecase{
		hospitalRepo: hospitalRepo,
	}
}

func (hu *hospitalUsecase) CreateHospital(ctx context.Context, hospital *models.Hospital) (*models.Hospital, error) {
	newHospital, err := hu.hospitalRepo.CreateHospital(ctx, hospital)
	if err != nil {
		return nil, err
	}
	return newHospital, nil
}

func (hu *hospitalUsecase) GetHospitalById(ctx context.Context, id string) (*models.Hospital, error) {
	hospital, err := hu.hospitalRepo.GetHospitalById(ctx, id)
	if err != nil {
		return nil, err
	}
	return hospital, nil
}
func (hu *hospitalUsecase) GetAllHospitals(ctx context.Context) ([]models.Hospital, error) {
	hospitals, err := hu.hospitalRepo.GetAllHospitals(ctx)
	if err != nil {
		return nil, err
	}
	return hospitals, nil
}
func (hu *hospitalUsecase) GetHospitalsByQuery(ctx context.Context, filter bson.M) ([]models.Hospital, error) {
	hospitals, err := hu.hospitalRepo.GetHospitalsByQuery(ctx, filter)
	if err != nil {
		return nil, err
	}
	return hospitals, nil
}
func (hu *hospitalUsecase) UpdateHospitalById(ctx context.Context, id string, updateQuery bson.M) (*models.Hospital, error) {
	updatedHospital, err := hu.hospitalRepo.UpdateHospitalById(ctx, id, updateQuery)
	if err != nil {
		return nil, err
	}
	return updatedHospital, nil
}
func (hu *hospitalUsecase) DeleteHospital(ctx context.Context, id string) (int64, error) {
	deletedCount, err := hu.hospitalRepo.DeleteHospital(ctx, id)
	if err != nil {
		return 0, err
	}
	return deletedCount, nil
}
