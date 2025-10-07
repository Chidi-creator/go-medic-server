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

type HospitalHandler interface {
	CreateHospital(w http.ResponseWriter, r *http.Request)
	GetHospitalById(w http.ResponseWriter, r *http.Request)
	GetAllHospitals(w http.ResponseWriter, r *http.Request)
	UpdateHospitalById(w http.ResponseWriter, r *http.Request)
	DeleteHospital(w http.ResponseWriter, r *http.Request)
}

type hospitalHandler struct {
	hu usecases.HospitalUsecase
}

func NewHospitalHandler(hu usecases.HospitalUsecase) HospitalHandler {
	return &hospitalHandler{
		hu: hu,
	}
}

func (h *hospitalHandler) CreateHospital(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var hospital models.Hospital
	err := json.NewDecoder(r.Body).Decode(&hospital)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
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

	managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
		Success: true,
		Message: "Successfully retrieved user",
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
