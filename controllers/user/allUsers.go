package user

import (
	"fmt"
	"net/http"
)

// GetAllUsers is a dummy endpoint to demonstrate authentication.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected endpoint")
}
