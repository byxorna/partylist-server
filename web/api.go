package web

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Error error `json:"error"`
}

// return an error as json
func ApiError(w http.ResponseWriter, e error) {
	jsonError := apiError{
		Error: e,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(jsonError); err != nil {
		panic(err)
	}
}
