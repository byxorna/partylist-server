package web

import (
	"log"
	"net/http"
	"time"
)

// wrap a http.Handler in a logging function
func AccessLogger(fn http.HandlerFunc, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn.ServeHTTP(w, r)
		log.Printf("[web] %s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}
