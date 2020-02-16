package routes

import (
	"go-rest-api-template/controllers/auth"

	"github.com/gorilla/mux"
)

// AuthRoutes contains all routes related to authentication.
func AuthRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/login", auth.Login).Methods("POST")
	subrouter.HandleFunc("/signup", auth.SignUp).Methods("POST")
}
