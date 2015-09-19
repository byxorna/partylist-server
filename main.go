package main

import (
	"flag"
	"fmt"
	log "github.com/golang/glog"
	"net/http"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/byxorna/partylist-server/web"
)

var (
	httpPort int
	dbPort   int
	dbHost   string
	dbName   string
	dbUser   string
	dbPass   string
)

func init() {
	flag.IntVar(&httpPort, "port", 8000, "HTTP port")
	flag.StringVar(&dbHost, "db-host", "localhost", "DB host")
	flag.IntVar(&dbPort, "db-port", 5432, "DB port")
	flag.StringVar(&dbName, "db-name", "partylist", "DB name")
	flag.StringVar(&dbUser, "db-user", "partylist", "DB user")
	flag.StringVar(&dbPass, "db-pass", "partylist", "DB password")
	flag.Parse()
}

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s port=%d host=%s sslmode=disable", dbUser, dbName, dbPort, dbHost))
	if err != nil {
		log.Fatal(err)
	}

	router := web.New(*db)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router); err != nil {
		log.Fatal(err)
	}
}
