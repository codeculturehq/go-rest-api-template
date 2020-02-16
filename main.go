package main

import (
	"fmt"
	"go-rest-api-template/initialization"
	"go-rest-api-template/keys"
	"go-rest-api-template/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

var port = fmt.Sprintf(":%s", keys.GetKeys().PORT)

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
