package handlers

import (
	"encoding/json"
	"github/Chidi-creator/go-medic-server/internal/managers"
	"github/Chidi-creator/go-medic-server/internal/middleware"
	"github/Chidi-creator/go-medic-server/internal/services"
	"github/Chidi-creator/go-medic-server/internal/usecases"
	"github/Chidi-creator/go-medic-server/internal/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandler struct {
	uc usecases.UserUseCase
}

func NewAuthHandler(uc usecases.UserUseCase) *AuthHandler {
	return &AuthHandler{
		uc: uc,
	}
}

func (ah *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var details utils.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "invalid request body: " + err.Error(),
		})
		return
	}

	// validate request
	validationErrs := utils.ValidateStruct(details)
	if validationErrs != "" {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   validationErrs,
		})
		return
	}
	// call usecase
	resp, err := ah.uc.LoginUser(ctx, &details)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "could not login user: " + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "User logged in successfully",
		Data:    resp,
	})

}
func (ah *AuthHandler) GenerateRefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.GetUserFromContext(ctx)
	if user == nil {
		managers.JSONresponse(w, http.StatusUnauthorized, utils.ApiResponse{
			Success: false,
			Error:   "User not authenticated",
		})
		return
	}

	userId, err := primitive.ObjectIDFromHex(user.UserID)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid user ID",
		})
		return
	}

	token, err := services.GenerateToken(userId, "refresh")
	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Could not generate refresh token: " + err.Error(),
		})
		return
	}
	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Refresh token generated successfully",
		Data:    token,
	})
}
