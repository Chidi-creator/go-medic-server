package handlers

import (
	"encoding/json"

	"github/Chidi-creator/go-medic-server/internal/managers"
	"github/Chidi-creator/go-medic-server/internal/models"
	"github/Chidi-creator/go-medic-server/internal/usecases"
	"github/Chidi-creator/go-medic-server/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type UserHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	UpdateUserById(w http.ResponseWriter, r *http.Request)
	DeleteUserById(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	uc usecases.UserUseCase
}

func NewUserHandler(uc usecases.UserUseCase) UserHandler {
	return &userHandler{
		uc: uc,
	}
}

func (uh *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid Request Body",
		})
		return

	}

	// validate request
	validationErrs := utils.ValidateStruct(user)
	if validationErrs != "nil" {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   validationErrs,
		})
		return
	}

	registeredUser, err := uh.uc.RegisterUser(ctx, &user)

	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Could not register user: " + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusCreated, utils.ApiResponse{
		Success: true,
		Message: "User Created Successfully",
		Data:    registeredUser,
	})

}

// <---- Get User By Id
func (uh *userHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)

	userId := params["id"]

	user, err := uh.uc.GetUserById(ctx, userId)
	if err != nil {
		managers.JSONresponse(w, http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Error:   "Could not find User: " + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})

}

// <----  Update user by Id
func (uh *userHandler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)

	userId := params["id"]

	user, err := uh.uc.GetUserById(ctx, userId)

	if err != nil {
		managers.JSONresponse(w, http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Error:   "Error finding user: " + err.Error(),
		})
		return
	}

	if user == nil {
		managers.JSONresponse(w, http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Error:   "User doesnt exist ",
		})
		return
	}

	var update bson.M

	err = json.NewDecoder(r.Body).Decode(&update)

	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid update payload",
		})
		return
	}

	err = uh.uc.UpdateUserById(ctx, userId, update)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Could not update user",
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "User updated successfully",
	})
}

func (uh *userHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)

	userId := params["id"]

	count, err := uh.uc.DeleteUserById(ctx, userId)

	if err != nil {
		managers.JSONresponse(w, http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Error:   "Could not delete user ",
		})
		return
	}

	if count == 0 {
		managers.JSONresponse(w, http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Error:   "No user got deleted",
		})
		return
	}

	managers.JSONresponse(w, http.StatusNoContent, utils.ApiResponse{
		Success: true,
		Message: "User deleted Successfully",
	})
}
