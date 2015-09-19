package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"ApiIndex", "GET", "/api/v1", ApiV1Index},
	Route{"CreatePlaylist", "POST", "/api/v1/playlist", ApiV1CreatePlaylist},
	Route{"ShowPlaylist", "GET", "/api/v1/playlist/{plid}", ApiV1GetPlaylist},
	Route{"DeletePlaylist", "DELETE", "/api/v1/playlist/{plid}", ApiV1DeletePlaylist},
	Route{"AddSong", "POST", "/api/v1/playlist/{plid}/add", ApiV1AddSong},
	Route{"DeleteSong", "DELETE", "/api/v1/playlist/{plid}/{sid}", ApiV1DeleteSong},
}

func New() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
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
