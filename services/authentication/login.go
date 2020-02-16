package authentication

import (
	"errors"
	"go-rest-api-template/models"
	"go-rest-api-template/services/user"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login checks if the user with the given email exists and verifies the given password.
func Login(userData models.User) (string, error) {
	user, err := user.FindOne(bson.M{"email": userData.Email})
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
