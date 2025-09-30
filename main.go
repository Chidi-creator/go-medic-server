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
	userCollection     = "users"
	hospitalCollection = "hospitals"
	doctorCollection   = "doctors"
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

	//initialising usecases
	userUsecase := usecases.NewUserUsecase(userRepo)
	doctorUsecase := usecases.NewDoctorUseCase(doctorRepo)

	//initializing handlers
	userHandler := handlers.NewUserHandler(userUsecase)
	doctorHandler := handlers.NewDoctorHandler(doctorUsecase)

	r := routes.NewRouter(userHandler, doctorHandler)
	r.SetUpRoutes()

	fmt.Printf("Server started on %v", config.AppConfig.Port)

	//initializing repositories

	http.ListenAndServe(":"+config.AppConfig.Port, r.R)

}
