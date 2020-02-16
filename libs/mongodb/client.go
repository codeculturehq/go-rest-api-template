package mongodb

import (
	"context"
	"go-rest-api-template/keys"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Database

// InitiateDatabase creates a MongoDB client and establishes a connection and assigns a pointer to the database to db.
func InitiateDatabase() {
	if db != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(keys.GetKeys().MONGO_URI),
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("❌", err.Error())
	}

	log.Println("✅ Connection to MongoDB established")

	db = client.Database(keys.GetKeys().MONGO_DB_NAME)
}

// GetClient returns a pointer to the MongoDB database that can be used throughout the application.
func GetClient() *mongo.Database {
	if db != nil {
		return db
	}

	InitiateDatabase()

	return db
}
