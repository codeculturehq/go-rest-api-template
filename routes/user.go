package routes

import (
	"go-rest-api-template/controllers/user"
	"go-rest-api-template/middlewares"

	"github.com/gorilla/mux"
)

// UserRoutes contains all routes related to user management.
func UserRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/all", user.GetAllUsers).Methods("GET")
	subrouter.Use(middlewares.VerifyJwt)
}
