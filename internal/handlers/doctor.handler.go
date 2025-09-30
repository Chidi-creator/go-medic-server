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

type DoctorHandler interface {
	CreateDoctor(w http.ResponseWriter, r *http.Request)
	FindDoctorById(w http.ResponseWriter, r *http.Request)
	GetDoctorsByHospitalId(w http.ResponseWriter, r *http.Request)
	UpdateDoctorById(w http.ResponseWriter, r *http.Request)
	DeleteDoctorByUserId(w http.ResponseWriter, r *http.Request)
}

type doctorHandler struct {
	doctorusecase usecases.DoctorUsecase
}

func NewDoctorHandler(du usecases.DoctorUsecase) DoctorHandler {
	return &doctorHandler{
		doctorusecase: du,
	}
}

func (dh *doctorHandler) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var doctor models.Doctor

	err := json.NewDecoder(r.Body).Decode(&doctor)

	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	newDoctor, err := dh.doctorusecase.CreateDoctor(ctx, &doctor)

	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Could not create doctor" + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Doctor successfully registered",
		Data:    newDoctor,
	})

}

func (dh *doctorHandler) FindDoctorById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	param := mux.Vars(r)

	id := param["id"]

	doctor, err := dh.doctorusecase.FindDoctorById(ctx, id)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Could not find doctor" + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Doctor successfully retrieved",
		Data:    doctor,
	})

}

func (dh *doctorHandler) GetDoctorsByHospitalId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)

	id := params["id"]

	doctors, err := dh.doctorusecase.GetDoctorsByHospitalId(ctx, id)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Could not fetch doctors" + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Doctor successfully retrieved",
		Data:    doctors,
	})

}

func (dh *doctorHandler) UpdateDoctorById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)

	id := params["id"]

	var update bson.M

	err := json.NewDecoder(r.Body).Decode(&update)

	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	err = dh.doctorusecase.UpdateDoctorById(ctx, id, update)

	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Failed to update doctor: " + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Doctor successfully updated",
	})
}

func (dh *doctorHandler) DeleteDoctorByUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)

	id := params["id"]

	err := dh.doctorusecase.DeleteDoctorByUserId(ctx, id)

	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Could not delete user from db" + err.Error(),
		})
		return
	}
	managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
		Success: false,
		Error:   "Doctor successfully deleted from db",
	})
}
