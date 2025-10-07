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

type AppointmentHandler interface {
	CreateAppointment(w http.ResponseWriter, r *http.Request)
	GetSingleAppointmentById(w http.ResponseWriter, r *http.Request)
	GetAppointmentsByDoctorId(w http.ResponseWriter, r *http.Request)
	GetAppointmentsByUserId(w http.ResponseWriter, r *http.Request)
	UpdateAppointmentById(w http.ResponseWriter, r *http.Request)
	DeleteAppointmentById(w http.ResponseWriter, r *http.Request)
}

type appointmentHandler struct {
	appointmentUsecase usecases.AppointmentUsecase
}

func NewAppointmentHandler(au usecases.AppointmentUsecase) AppointmentHandler {
	return &appointmentHandler{
		appointmentUsecase: au,
	}

}

func (a *appointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var appointment models.Appointment

	err := json.NewDecoder(r.Body).Decode(&appointment)

	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid request body: " + err.Error(),
		})
		return
	}

	newAppointment, err := a.appointmentUsecase.CreateAppointment(ctx, &appointment)

	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Internal Server Error: " + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusCreated, utils.ApiResponse{
		Success: true,
		Message: "Appointment created successfully",
		Data:    newAppointment,
	})
}

func (a *appointmentHandler) GetSingleAppointmentById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	id := params["id"]

	appointment, err := a.appointmentUsecase.GetSingleAppointmentById(ctx, id)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Could not fetch appointment: " + err.Error(),
		})
		return
	}

	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    appointment,
	})
}

func (a *appointmentHandler) GetAppointmentsByDoctorId(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	appointments, err := a.appointmentUsecase.GetAppointmentsByDoctorId(ctx, id)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Could not fetch appointments: " + err.Error(),
		})
		return
	}
	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Doctor's Appointments successfully retrieved",
		Data:    appointments,
	})
}

func (a *appointmentHandler) GetAppointmentsByUserId(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	params := mux.Vars(r)

	id := params["id"]
	appointments, err := a.appointmentUsecase.GetAppointmentsByUserId(ctx, id)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Could not fetch appointments: " + err.Error(),
		})
		return
	}
	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "User's Appointments successfully retrieved",
		Data:    appointments,
	})
}

func (a *appointmentHandler) UpdateAppointmentById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	var updateData bson.M
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		managers.JSONresponse(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid request body: " + err.Error(),
		})
		return
	}
	updatedAppointment, err := a.appointmentUsecase.UpdateAppointmentById(ctx, id, updateData)
	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Could not update appointment: " + err.Error(),
		})
		return
	}
	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Appointment updated successfully",
		Data:   updatedAppointment,
	})
}

func (a *appointmentHandler) DeleteAppointmentById(w http.ResponseWriter, r *http.Request) {	
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	deletedCount, err := a.appointmentUsecase.DeleteAppointmentById(ctx, id)
	if err != nil {
		managers.JSONresponse(w, http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Error:   "Could not delete appointment: " + err.Error(),
		})
		return
	}
	if deletedCount == 0 {
		managers.JSONresponse(w, http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Error:   "Appointment not found",
		})
		return
	}
	managers.JSONresponse(w, http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Appointment deleted successfully",
	})
}
