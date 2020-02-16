package user

import (
	"context"
	"errors"
	"go-rest-api-template/libs/mongodb"
	"go-rest-api-template/models"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create stores a new user in the database.
func Create(newUser models.User) (models.MongoDbDocument, error) {
	var createdDocument models.MongoDbDocument

	collection := mongodb.GetClient().Collection("users")
	result, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Println("Failed to insert new user with error:", err.Error())
		return createdDocument, errors.New("An error occured while creating the user")
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		createdDocument.ID = &objectID
	}

	return createdDocument, nil
}
