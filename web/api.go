package web

import (
	"encoding/json"
	"net/http"

	log "github.com/golang/glog"
)

type apiError struct {
	Error error `json:"error"`
}

// return an error as json
func ApiError(w http.ResponseWriter, statusCode int, clientError error, realError error) {
	w.WriteHeader(statusCode)
	jsonError := apiError{
		Error: clientError,
	}
	log.Error(realError)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(jsonError); err != nil {
		panic(err)
	}
}
