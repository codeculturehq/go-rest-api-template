package main

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/initialization"
	"rest_api/keys"
	"rest_api/routes"

	"github.com/gorilla/handlers"
)

var port = fmt.Sprintf(":%s", keys.GetKeys().PORT)

/*
	Initializes environment variables and establishes
	a connection to MongoDB.
*/
func init() {
	initialization.InitEnv()
	initialization.InitDatabase()
}

func main() {
	log.Println("🆙 Server listening on port", port)
	http.ListenAndServe(port, handlers.CORS()(routes.InitRouters()))
}
