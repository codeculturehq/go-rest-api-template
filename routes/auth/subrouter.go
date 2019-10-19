package authroutes

import (
	"github.com/gorilla/mux"
)

/*
	Registers routes to the "/auth" subrouter.
*/
func InitRouter(subrouter *mux.Router) {
	subrouter.HandleFunc("/login", Login).Methods("POST")
	subrouter.HandleFunc("/signup", SignUp).Methods("POST")
}
