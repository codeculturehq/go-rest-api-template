package usermanagementroutes

import (
	"rest_api/middlewares"

	"github.com/gorilla/mux"
)

/*
	Registers routes to the "/usermanagement" subrouter and
	adds the middleware to verify JWT's to the subrouter.
*/
func InitRouter(subrouter *mux.Router) {
	subrouter.HandleFunc("/allusers", GetAllUsers).Methods("GET")
	subrouter.Use(middlewares.VerifyJwt)
}
