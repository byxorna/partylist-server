package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/byxorna/partylist-server/models"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	SongEnqueueError = errors.New("Unable to enqueue song")
	SongDequeueError = errors.New("Unable to dequeue song")
	SongsGetError    = errors.New("Unable to get songs")
)

func ApiV1GetSongsForPlaylist(w http.ResponseWriter, r *http.Request) {
	requestedPlaylistId := mux.Vars(r)["plid"]
	//TODO we should validate this playlist is accessible to the requestor

	//check if the requested plid is actually a contributor id and find the real playlist
	masterPlaylistId, err := redisClient.HGet("contributor_to_playlist", requestedPlaylistId).Result()
	if err != nil {
		ApiError(w, 500, SongsGetError, fmt.Errorf("Unable to lookup master playlist id from contributor id %s: %s", requestedPlaylistId, err))
		return
	}
	if masterPlaylistId == "" {
		masterPlaylistId = requestedPlaylistId
	}

	// return all songs attached to a playlist
	nsongs, err := redisClient.LLen("songs:" + masterPlaylistId).Result()
	if err != nil {
		ApiError(w, http.StatusInternalServerError, SongsGetError, fmt.Errorf("Unable to query for number of songs in playlist %s: %s", masterPlaylistId, err))
		return
	}

	// fetch all songs by paginating
	panic("fuck")
}

func ApiV1EnqueueSong(w http.ResponseWriter, r *http.Request) {
	//TODO add song to playlist
}

func ApiV1DequeueSong(w http.ResponseWriter, r *http.Request) {
	//TODO delete song from playlist
}
