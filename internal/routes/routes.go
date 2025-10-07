package routes

import (
	"github/Chidi-creator/go-medic-server/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	R                  *mux.Router
	UserHandler        handlers.UserHandler
	DoctorHandler      handlers.DoctorHandler
	AppointmentHandler handlers.AppointmentHandler
}

func NewRouter(h handlers.UserHandler,
	d handlers.DoctorHandler,
	a handlers.AppointmentHandler,
) *Router {
	return &Router{
		R:                  mux.NewRouter(),
		UserHandler:        h,
		DoctorHandler:      d,
		AppointmentHandler: a,
	}
}

func (r *Router) SetUpRoutes() {
	log.Println("Setting up routes")
	r.R.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Chidi Medic Server Up and Running<h1>"))
	})

	//user routes

	userRouter := r.R.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("", r.UserHandler.RegisterUser).Methods("POST")
	userRouter.HandleFunc("/login", r.UserHandler.LoginUser).Methods("POST")
	userRouter.HandleFunc("/{id}", r.UserHandler.GetUserById).Methods("GET")

	doctorRouter := r.R.PathPrefix("/doctors").Subrouter()

	doctorRouter.HandleFunc("", r.DoctorHandler.CreateDoctor).Methods("POST")
	doctorRouter.HandleFunc("/{id}", r.DoctorHandler.FindDoctorById).Methods("GET")
	doctorRouter.HandleFunc("/{id}", r.DoctorHandler.GetDoctorsByHospitalId).Methods("GET")
	doctorRouter.HandleFunc("/{id}", r.DoctorHandler.UpdateDoctorById).Methods("PATCH")
	doctorRouter.HandleFunc("/{id}", r.DoctorHandler.DeleteDoctorByUserId).Methods("DELETE")

	appointmentRouter := r.R.PathPrefix("/appointments").Subrouter()

	appointmentRouter.HandleFunc("", r.AppointmentHandler.CreateAppointment).Methods("POST")
	appointmentRouter.HandleFunc("/{id}", r.AppointmentHandler.GetSingleAppointmentById).Methods("GET")
	appointmentRouter.HandleFunc("/user/{id}", r.AppointmentHandler.GetAppointmentsByUserId).Methods("GET")
	appointmentRouter.HandleFunc("/doctor/{id}", r.AppointmentHandler.GetAppointmentsByDoctorId).Methods("GET")
	appointmentRouter.HandleFunc("/{id}", r.AppointmentHandler.UpdateAppointmentById).Methods("PATCH")
	appointmentRouter.HandleFunc("/{id}", r.AppointmentHandler.DeleteAppointmentById).Methods("DELETE")

}
