package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	SongEnqueueError = errors.New("Unable to enqueue song")
	SongDequeueError = errors.New("Unable to dequeue song")
	SongsGetError    = errors.New("Unable to get songs")
)

func ApiV1GetSongsForPlaylist(c *gin.Context) {
	requestedPlaylistId := c.Param("plid")
	//TODO we should validate this playlist is accessible to the requestor

	//check if the requested plid is actually a contributor id and find the real playlist
	masterPlaylistId, err := redisClient.HGet("contributor_to_playlist", requestedPlaylistId).Result()
	if err != nil {
		ApiError(c, 500, SongsGetError, fmt.Errorf("Unable to lookup master playlist id from contributor id %s: %s", requestedPlaylistId, err))
		return
	}
	if masterPlaylistId == "" {
		masterPlaylistId = requestedPlaylistId
	}

	// return all songs attached to a playlist
	nsongs, err := redisClient.LLen("songs:" + masterPlaylistId).Result()
	if err != nil {
		ApiError(c, http.StatusInternalServerError, SongsGetError, fmt.Errorf("Unable to query for number of songs in playlist %s: %s", masterPlaylistId, err))
		return
	}
	nsongs = nsongs
	// fetch all songs by paginating
	panic("fuck FIXME")
}

func ApiV1EnqueueSong(c *gin.Context) {
	//TODO add song to playlist
}

func ApiV1DequeueSong(c *gin.Context) {
	//TODO delete song from playlist
}
