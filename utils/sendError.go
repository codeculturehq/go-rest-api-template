package utils

import (
	"encoding/json"
	"go-rest-api-template/models"
	"net/http"
)

// SendError formats an error response.
func SendError(w http.ResponseWriter, error models.Error) {
	w.WriteHeader(error.Code)
	json.NewEncoder(w).Encode(error)
}
