package main

import (
	"fmt"
	"github/Chidi-creator/go-medic-server/config"
	"github/Chidi-creator/go-medic-server/internal/handlers"
	"github/Chidi-creator/go-medic-server/internal/mongo"
	"github/Chidi-creator/go-medic-server/internal/repositories"
	"github/Chidi-creator/go-medic-server/internal/routes"
	"github/Chidi-creator/go-medic-server/internal/usecases"
	_ "github/Chidi-creator/go-medic-server/internal/utils"
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
	
	//create indexes after successful mongo connecttion
	mongo.CreateIndexes(client.Client, config.AppConfig.DB_NAME)

	//initialising repositories
	userRepo := repositories.NewUserRepository(client.Client, config.AppConfig.DB_NAME, userCollection)
	doctorRepo := repositories.NewDoctorRepository(client.Client, config.AppConfig.DB_NAME, doctorCollection)
	hospitalRepo := repositories.NewHospitalRepository(client.Client, config.AppConfig.DB_NAME, hospitalCollection)
	appointmentRepo := repositories.NewAppointmentRepository(client.Client, config.AppConfig.DB_NAME, appointmentCollection)

	//initialising usecases
	userUsecase := usecases.NewUserUsecase(userRepo)
	doctorUsecase := usecases.NewDoctorUseCase(doctorRepo)
	hospitalUsecase := usecases.NewHospitalUseCase(hospitalRepo)
	appointmentUsecase := usecases.NewAppointmentUsecase(appointmentRepo)

	//initializing handlers
	userHandler := handlers.NewUserHandler(userUsecase)
	doctorHandler := handlers.NewDoctorHandler(doctorUsecase)
	hospitalHandler := handlers.NewHospitalHandler(hospitalUsecase, userUsecase)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentUsecase)
	authHandler := handlers.NewAuthHandler(userUsecase)

	r := routes.NewRouter(userHandler, doctorHandler, hospitalHandler, appointmentHandler, authHandler)
	r.SetUpRoutes()

	fmt.Printf("Server started on %v", config.AppConfig.Port)

	http.ListenAndServe(":"+config.AppConfig.Port, r.R)

}
