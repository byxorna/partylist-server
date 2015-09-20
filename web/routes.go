package web

import (
	"github.com/gin-gonic/gin"
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

func New(redisclient *redis.Client) *gin.Engine {
	//TODO how can you pass a variable into a gin Engine's context? This is dirty af
	redisClient = redisclient

	router := gin.Default()
	router.GET("/", Index)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", ApiV1Index)

		v1.POST("/playlist", ApiV1CreatePlaylist)
		v1.GET("/playlist/:plid", ApiV1GetPlaylist)
		v1.DELETE("/playlist/:plid", ApiV1DeletePlaylist)

		v1.GET("/playlist/:plid/songs", ApiV1GetSongsForPlaylist)
		v1.POST("/playlist/:plid/enqueue", ApiV1EnqueueSong)
		v1.DELETE("/playlist/:plid/:sid", ApiV1DequeueSong)
	}

	return router
}
