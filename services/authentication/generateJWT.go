package authentication

import (
	"go-rest-api-template/keys"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJwt generates a JWT for a given user.
func GenerateJwt(userID string) (string, error) {
	timestamp := time.Now().Unix()
	expiresAt := timestamp + 600

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"iat": timestamp,
		"exp": expiresAt,
	})

	signedToken, err := token.SignedString([]byte(keys.GetKeys().SECRET))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
