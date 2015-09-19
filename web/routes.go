package web

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type HandlerFuncWithDb func(w http.ResponseWriter, r *http.Request, db *sql.DB)

type Routes []Route

//TODO i dont like this. how can we pass the DB context into handlers while satisfying http.HandlerFunc interface
var (
	db sql.DB
)

func New(database sql.DB) *mux.Router {
	db = database

	router := mux.NewRouter().StrictSlash(true)
	routes := Routes{
		Route{"Index", "GET", "/", Index},
		Route{"ApiIndex", "GET", "/api/v1", ApiV1Index},
		Route{"CreatePlaylist", "POST", "/api/v1/playlist", ApiV1CreatePlaylist},
		Route{"ShowPlaylist", "GET", "/api/v1/playlist/{plid}", ApiV1GetPlaylist},
		Route{"DeletePlaylist", "DELETE", "/api/v1/playlist/{plid}", ApiV1DeletePlaylist},
		Route{"AddSong", "POST", "/api/v1/playlist/{plid}/add", ApiV1AddSong},
		Route{"DeleteSong", "DELETE", "/api/v1/playlist/{plid}/{sid}", ApiV1DeleteSong},
	}

	for _, route := range routes {
		loggingHandler := AccessLogger(route.HandlerFunc, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(loggingHandler)
	}

	return router
}
