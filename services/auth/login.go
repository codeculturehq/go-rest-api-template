package authservices

import (
	"errors"
	"rest_api/models"
	userservices "rest_api/services/users"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

/*
	Checks if the user with the given email exists and
	verifies the given password.
*/
func Login(userData models.User) (string, error) {
	user, err := userservices.FindOne(bson.M{"email": userData.Email})

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		return "", errors.New("Incorrect password")
	}

	jwt, err := GenerateJwt(string(user.ID.Hex()))

	if err != nil {
		return "", err
	}

	return jwt, nil
}
