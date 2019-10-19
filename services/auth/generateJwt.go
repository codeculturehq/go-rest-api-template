package authservices

import (
	"rest_api/keys"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
	Generates and returns a JWT.
*/
func GenerateJwt(userId string) (string, error) {
	timestamp := time.Now().Unix()
	expiresAt := timestamp + 600

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"iat": timestamp,
		"exp": expiresAt,
	})

	signedToken, err := token.SignedString([]byte(keys.GetKeys().SECRET))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
