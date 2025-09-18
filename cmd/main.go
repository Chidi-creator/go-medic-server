package main

import (
	"fmt"
	"github/Chidi-creator/go-medic-server/config"
	"github/Chidi-creator/go-medic-server/internal/mongo"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mongo.ConnectMongo()

	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Chidi Medic Server Up and Running<h1>"))
	})

	fmt.Printf("Server started on %v", config.AppConfig.Port)

	http.ListenAndServe(":8080", r)

}
