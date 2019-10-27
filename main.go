package main

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/initialization"
	"rest_api/keys"
	"rest_api/routes"
	"time"

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
	srv := &http.Server{
		Addr:         port,
		Handler:      handlers.CORS()(routes.InitRouters()),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	log.Println("🆙 Server listening on port", port)
	srv.ListenAndServe()
}
