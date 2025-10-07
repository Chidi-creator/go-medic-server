package main

import (
	"fmt"
	"github/Chidi-creator/go-medic-server/config"
	"github/Chidi-creator/go-medic-server/internal/handlers"
	"github/Chidi-creator/go-medic-server/internal/mongo"
	"github/Chidi-creator/go-medic-server/internal/repositories"
	"github/Chidi-creator/go-medic-server/internal/routes"
	"github/Chidi-creator/go-medic-server/internal/usecases"
	"log"

	"net/http"
)

var (
	userCollection        = "users"
	hospitalCollection    = "hospitals"
	doctorCollection      = "doctors"
	appointmentCollection = "appointments"
)

func main() {
	//connect to mongoDB

	client, err := mongo.NewClient(config.AppConfig.Mongo_URI, config.AppConfig.DB_NAME)
	if err != nil {
		log.Fatal("Could not connect to Mongo DB")
	}

	//initialising repositories
	userRepo := repositories.NewUserRepository(client.Client, config.AppConfig.DB_NAME, userCollection)
	doctorRepo := repositories.NewDoctorRepository(client.Client, config.AppConfig.DB_NAME, doctorCollection)
	appointmentRepo := repositories.NewAppointmentRepository(client.Client, config.AppConfig.DB_NAME, appointmentCollection)

	//initialising usecases
	userUsecase := usecases.NewUserUsecase(userRepo)
	doctorUsecase := usecases.NewDoctorUseCase(doctorRepo)
	appointmentUsecase := usecases.NewAppointmentUsecase(appointmentRepo)

	//initializing handlers
	userHandler := handlers.NewUserHandler(userUsecase)
	doctorHandler := handlers.NewDoctorHandler(doctorUsecase)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentUsecase)

	r := routes.NewRouter(userHandler, doctorHandler, appointmentHandler)
	r.SetUpRoutes()

	fmt.Printf("Server started on %v", config.AppConfig.Port)


	http.ListenAndServe(":"+config.AppConfig.Port, r.R)

}
