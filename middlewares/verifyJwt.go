package middlewares

import (
	"fmt"
	"net/http"
	"rest_api/keys"
	"rest_api/models"
	"rest_api/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

/*
	Middleware to verify JWT's sent as Authorization header.
*/
func VerifyJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var errorResponse models.Error

		authenticationHeader := r.Header.Get("Authorization")
		splitHeader := strings.Split(authenticationHeader, " ")

		if len(splitHeader) != 2 {
			errorResponse.Code = http.StatusBadRequest
			errorResponse.Message = "Incorrect header"

			utils.SendError(w, errorResponse)

			return
		}

		token, err := jwt.Parse(splitHeader[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("An error occured while verifying the token")
			}

			return []byte(keys.GetKeys().SECRET), nil
		})

		if err != nil || !token.Valid {
			errorResponse.Code = http.StatusUnauthorized
			errorResponse.Message = "Invalid JWT"

			utils.SendError(w, errorResponse)

			return
		}

		next.ServeHTTP(w, r)
	})
}
