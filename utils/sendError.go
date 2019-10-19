package utils

import (
	"encoding/json"
	"net/http"
	"rest_api/models"
)

func SendError(w http.ResponseWriter, error models.Error) {
	w.WriteHeader(error.Code)
	json.NewEncoder(w).Encode(error)
}
