package main

import (
	log "github.com/golang/glog"
	"net/http"

	"github.com/byxorna/partylist-server/web"
)

func main() {
	router := web.New()
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
