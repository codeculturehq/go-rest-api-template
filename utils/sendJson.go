package utils

import (
	"encoding/json"
	"net/http"
)

// SendJSON responds to a request with JSON.
func SendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
