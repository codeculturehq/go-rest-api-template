package routes

import (
	authroutes "rest_api/routes/auth"
	usermanagementroutes "rest_api/routes/userManagement"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

/*
	Creates and returns a subrouter for the given route.
*/
func createSubrouter(prefix string) *mux.Router {
	return Router.PathPrefix(prefix).Subrouter()
}

/*
	Adds all subrouters to the router instance.
*/
func InitRouters() *mux.Router {
	authroutes.InitRouter(createSubrouter("/auth"))
	usermanagementroutes.InitRouter(createSubrouter("/usermanagement"))

	return Router
}
