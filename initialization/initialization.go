package initialization

import (
	"go-rest-api-template/keys"
	"go-rest-api-template/libs/mongodb"
	"log"
)

// InitEnv loads all environment variables.
func InitEnv() {
	keys.GetKeys()
	log.Println("✅ environment variables initialized successfully")
}

// InitDatabase creates a connection to the database.
func InitDatabase() {
	mongodb.InitiateDatabase()
}
