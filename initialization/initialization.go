package initialization

import (
	"log"
	"rest_api/keys"
	"rest_api/libs/mongodb"
)

func InitEnv() {
	keys.GetKeys()

	log.Println("✅ environment variables initialized successfully")
}

func InitDatabase() {
	mongodb.InitiateDatabase()
}
