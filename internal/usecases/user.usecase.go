package usecases

import (
	"context"
	"fmt"

	"github/Chidi-creator/go-medic-server/internal/models"
	"github/Chidi-creator/go-medic-server/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUsersByQuery(ctz context.Context, filter bson.M) ([]models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	UpdateUserById(ctx context.Context, id string, updateQuery bson.M) error
	DeleteUserById(ctx context.Context, id string) (int64, error)
}

type userUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	savedUser, err := uc.userRepo.RegisterUser(ctx, user)

	if err != nil {
		return nil, err
	}
	return savedUser, nil

}

func (uc *userUseCase) GetUsersByQuery(ctx context.Context, filter bson.M) ([]models.User, error) {

	users, err := uc.userRepo.GetUsersByQuery(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not fetch users: %w", err)
	}
	return users, nil

}

func (uc *userUseCase) GetUserById(ctx context.Context, id string) (*models.User, error) {
	user, err := uc.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not fetch user by id: %w", err)
	}
	return user, nil
}

func (uc *userUseCase) UpdateUserById(ctx context.Context, id string, updateQuery bson.M) error {
	err := uc.userRepo.UpdateUserById(ctx, id, updateQuery)

	if err != nil {
		return fmt.Errorf("could not update user by id: %w", err)
	}
	return nil
}

func (uc *userUseCase) DeleteUserById(ctx context.Context, id string) (int64, error) {
	count, err := uc.userRepo.DeleteUserById(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("could not update user by id: %w", err)
	}
	return count, nil
}

