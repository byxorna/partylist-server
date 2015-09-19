package web

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ApiV1Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API v1 Index")
}
