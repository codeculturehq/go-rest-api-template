package routes

import "github.com/gorilla/mux"

var router = mux.NewRouter()

func createSubrouter(prefix string) *mux.Router {
	return router.PathPrefix(prefix).Subrouter()
}

// InitRouters adds all subrouters to the router instance.
func InitRouters() *mux.Router {
	AuthRoutes((createSubrouter("/auth")))
	UserRoutes((createSubrouter("/user")))

	return router
}
