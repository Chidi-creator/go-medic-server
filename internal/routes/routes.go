package routes

import (
	"github/Chidi-creator/go-medic-server/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	R          *mux.Router
	UserHandler handlers.UserHandler
}

func NewRouter(h handlers.UserHandler) *Router {
	return &Router{
		R:           mux.NewRouter(),
		UserHandler: h,
	}
}

func (r *Router) SetUpRoutes() {
	log.Println("Setting up routes")
	r.R.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Chidi Medic Server Up and Running<h1>"))
	})

	//user routes
	r.R.HandleFunc("/users", r.UserHandler.RegisterUser).Methods("POST")

}
