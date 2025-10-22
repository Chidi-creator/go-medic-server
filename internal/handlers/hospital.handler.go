package handlers

import (
	"encoding/json"
	"github/Chidi-creator/go-medic-server/internal/managers"
	"github/Chidi-creator/go-medic-server/internal/middleware"
	"github/Chidi-creator/go-medic-server/internal/models"
	"github/Chidi-creator/go-medic-server/internal/usecases"
	"github/Chidi-creator/go-medic-server/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HospitalHandler interface {
	CreateHospital(w http.ResponseWriter, r *http.Request)
	GetHospitalById(w http.ResponseWriter, r *http.Request)
	GetAllHospitals(w http.ResponseWriter, r *http.Request)
	UpdateHospitalById(w http.ResponseWriter, r *http.Request)
	DeleteHospital(w http.ResponseWriter, r *http.Request)
}

type hospitalHandler struct {
	hu usecases.HospitalUsecase
	uu usecases.UserUseCase
}

func NewHospitalHandler(hu usecases.HospitalUsecase, uu usecases.UserUseCase) HospitalHandler {
	return &hospitalHandler{
		hu: hu,
		uu: uu,
	}
}

func (h *hospitalHandler) CreateHospital(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		managers.JSONresponse(w, http.StatusUnauthorized, utils.ApiResponse{
			Success: false,
			Error:   "User not authenticated",
		})
		return
	}

	var hospital models.Hospital
	if err := json.NewDecoder(r.Body).Decode(&hospital); err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid request body: " + err.Error(),
		})
		return
	}

	// Set UserID from authenticated user (this prevents clients from spoofing the UserID)
	userId, err := primitive.ObjectIDFromHex(user.UserID)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid user ID format",
		})
		return
	}
	hospital.UserID = userId

	// Validate the struct
	validationErrs := utils.ValidateStruct(hospital)
	if validationErrs != "" {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Validation failed: " + validationErrs,
		})
		return
	}

	newHospital, err := h.hu.CreateHospital(ctx, &hospital)
	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}
	updateQuery := bson.M{
		"$addToSet": bson.M{"roles": utils.HOSPITAL},
	}

	err = h.uu.UpdateUserById(ctx, user.UserID, updateQuery)
	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Could not update user roles: " + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusCreated, utils.ApiResponse{
		Success: true,
		Message: "Hospital Successfully created",
		Data:    newHospital,
	})

}

func (h *hospitalHandler) GetHospitalById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)

	id := params["id"]

	hospital, err := h.hu.GetHospitalById(ctx, id)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Could no " + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Successfully retrieved hospital",
		Data:    hospital,
	})

}
func (h *hospitalHandler) GetAllHospitals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hospitals, err := h.hu.GetAllHospitals(ctx)
	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Could not retrieve hospitals: " + err.Error(),
		})
		return
	}
	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Hospitals retrieved successfully",
		Data:    hospitals,
	})
}

func (h *hospitalHandler) UpdateHospitalById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	id := params["id"]
	var update bson.M

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid request " + err.Error(),
		})
		return
	}

	updatedHospital, err := h.hu.UpdateHospitalById(ctx, id, update)
	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Could not update hospital: " + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Hospital uupdated successfully",
		Data:    updatedHospital,
	})
}

func (h *hospitalHandler) DeleteHospital(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	deletedCount, err := h.hu.DeleteHospital(ctx, id)
	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Could not delete hospital: " + err.Error(),
		})
		return
	}

	if deletedCount == 0 {
		managers.JSONresponse(w, http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Error:   "Hospital not found",
		})
		return
	}
	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Hospital deleted successfully",
	})
}
