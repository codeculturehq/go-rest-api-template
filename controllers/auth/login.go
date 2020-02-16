package auth

import (
	"encoding/json"
	"go-rest-api-template/models"
	"go-rest-api-template/services/authentication"
	"go-rest-api-template/utils"
	"net/http"
)

// Login verifies credentials and creates a JWT.
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var jwt models.Jwt
	var errorResponse models.Error

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "There was an error while parsing the request body"

		utils.SendError(w, errorResponse)

		return
	}

	if user.Email == "" {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "Email address was not provided"

		utils.SendError(w, errorResponse)

		return
	}

	if user.Password == "" {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "Password was not provided"

		utils.SendError(w, errorResponse)

		return
	}

	token, err := authentication.Login(user)
	if err != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "Could not login with credentials provided"

		utils.SendError(w, errorResponse)

		return
	}

	jwt.Token = token

	utils.SendJSON(w, jwt)
}
