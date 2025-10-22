package usecases

import (
	"context"
	"fmt"

	"github/Chidi-creator/go-medic-server/internal/models"
	"github/Chidi-creator/go-medic-server/internal/repositories"
	"github/Chidi-creator/go-medic-server/internal/services"
	"github/Chidi-creator/go-medic-server/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUsersByQuery(ctz context.Context, filter bson.M) ([]models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	UpdateUserById(ctx context.Context, id string, updateQuery bson.M) error
	DeleteUserById(ctx context.Context, id string) (int64, error)
	LoginUser(ctx context.Context, details *utils.LoginRequest) (map[string]interface{}, error)
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

	// assign a default role of customer if no role is provided
	if len(user.Roles) == 0 {
		user.Roles = append(user.Roles, utils.CUSTOMER)
	}

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

func (uc *userUseCase) LoginUser(ctx context.Context, details *utils.LoginRequest) (map[string]interface{}, error) {

	filter := bson.M{"email": details.Email}

	users, err := uc.userRepo.GetUsersByQuery(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	if len(users) == 0 || users == nil {
		return nil, fmt.Errorf("user with email doesn't exist")
	}

	user := users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(details.Password))
	if err != nil {
		return nil, fmt.Errorf("check email and password")
	}

	//generate token
	token, err := services.GenerateToken(user.ID, "access")

	if err != nil {
		return nil, fmt.Errorf("error generating user token: %w", err)
	}

	return map[string]interface{}{
		"user":  user,
		"token": token,
	}, nil

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
