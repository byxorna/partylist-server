package web

import (
	"github.com/gorilla/mux"
	"gopkg.in/redis.v3"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var (
	redisClient *redis.Client
)

func New(redisclient *redis.Client) *mux.Router {
	redisClient = redisclient

	router := mux.NewRouter().StrictSlash(true)
	routes := Routes{
		Route{"Index", "GET", "/", Index},
		Route{"ApiIndex", "GET", "/api/v1", ApiV1Index},
		Route{"CreatePlaylist", "POST", "/api/v1/playlist", ApiV1CreatePlaylist},
		Route{"ShowPlaylist", "GET", "/api/v1/playlist/{plid}", ApiV1GetPlaylist},
		Route{"DeletePlaylist", "DELETE", "/api/v1/playlist/{plid}", ApiV1DeletePlaylist},
		Route{"GetSongs", "GET", "/api/v1/playlist/{plid}/songs", ApiV1GetSongsForPlaylist},
		Route{"EnqueueSong", "POST", "/api/v1/playlist/{plid}/add", ApiV1EnqueueSong},
		Route{"DequeueSong", "DELETE", "/api/v1/playlist/{plid}/{sid}", ApiV1DequeueSong},
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
