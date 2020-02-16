package authentication

import (
	"errors"
	"go-rest-api-template/models"
	"go-rest-api-template/services/user"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// SignUp checks if the given email address was used before and stores the new user in the database with the hashed password.
func SignUp(newUser models.User) (string, error) {
	userExists, err := user.CheckExistence(bson.M{"email": newUser.Email})
	if err != nil {
		return "", err
	}

	if userExists {
		return "", errors.New("User with the given email address already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)
	if err != nil {
		return "", err
	}

	newUser.Password = string(hash)

	if _, err := user.Create(newUser); err != nil {
		return "", err
	}

	jwt, err := GenerateJwt(string(newUser.ID.Hex()))
	if err != nil {
		return "", err
	}

	return jwt, nil
}
