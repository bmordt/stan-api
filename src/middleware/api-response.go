package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"
)

//Used to allow nice json response format for errors
type ApiError struct{}

var error = ApiError{}

func ModelResponse(w http.ResponseWriter, responseCode int, responseBody interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	json.NewEncoder(w).Encode(responseBody)
	return
}

func (e ApiError) ApiError(w http.ResponseWriter, status int, message string) {
	error := make(map[string]string)

	error["error"] = message
	error["Status"] = strconv.Itoa(status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}
