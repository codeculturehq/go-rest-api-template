package authroutes

import (
	"encoding/json"
	"net/http"
	"rest_api/models"
	authservices "rest_api/services/auth"
	"rest_api/utils"
)

/*
	Handler for the "/signup" route.
*/
func SignUp(w http.ResponseWriter, r *http.Request) {
	var jwt models.Jwt
	var user models.User
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

	if user.FirstName == "" {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "First name was not provided"

		utils.SendError(w, errorResponse)

		return
	}

	if user.LastName == "" {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "Last name was not provided"

		utils.SendError(w, errorResponse)

		return
	}

	if user.Password == "" {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "Password was not provided"

		utils.SendError(w, errorResponse)

		return
	}

	token, err := authservices.SignUp(user)

	if err != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = err.Error()

		utils.SendError(w, errorResponse)

		return
	}

	jwt.Token = token

	utils.SendJson(w, jwt)
}
