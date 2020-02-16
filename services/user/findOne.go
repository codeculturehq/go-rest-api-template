package user

import (
	"context"
	"go-rest-api-template/libs/mongodb"
	"go-rest-api-template/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// FindOne returns one document of the users collection that matches the given filter.
func FindOne(filter bson.M) (models.User, error) {
	var user models.User

	collection := mongodb.GetClient().Collection("users")
	result := collection.FindOne(context.TODO(), filter)
	if result.Err() != nil {
		return user, result.Err()
	}

	if err := result.Decode(&user); err != nil {
		log.Println("Failed to decode user with error:", err.Error())
		return user, err
	}

	return user, nil
}
