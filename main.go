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
	userCollection = "users"
)

func main() {
	//connect to mongoDB

	client, err := mongo.NewClient(config.AppConfig.Mongo_URI, config.AppConfig.DB_NAME)
	if err != nil {
		log.Fatal("Could not connect to Mongo DB")
	}

	//initialising repositories
	userRepo := repositories.NewUserRepository(client.Client, config.AppConfig.DB_NAME, userCollection)

	//initialising usecases
	userUsecase := usecases.NewUserUsecase(userRepo)

	//initializing handlers
	userHandler := handlers.NewUserHandler(userUsecase)

	r := routes.NewRouter(userHandler)
	r.SetUpRoutes()

	fmt.Printf("Server started on %v", config.AppConfig.Port)

	//initializing repositories

	http.ListenAndServe(":" + config.AppConfig.Port, r.R)

}
