package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/byxorna/partylist-server/models"
	"github.com/gin-gonic/gin"
)

var (
	SongEnqueueError = errors.New("Unable to enqueue song")
	SongDequeueError = errors.New("Unable to dequeue song")
	SongsGetError    = errors.New("Unable to get songs")
)

const (
	SONG_FETCH_BATCH_SIZE = 50
)

func ApiV1GetSongTypes(c *gin.Context) {
	c.JSON(200, gin.H{
		"supported_types": []gin.H{
			gin.H{"typeid": models.SongTypeYoutube, "name": "YouTube", "icon_url": "xxx"},
			gin.H{"typeid": models.SongTypeSpotify, "name": "Spotify", "icon_url": "xxx"},
			gin.H{"typeid": models.SongTypeSoundcloud, "name": "Soundcloud", "icon_url": "xxx"},
		},
	})
}

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
	songsKey := "songs:" + masterPlaylistId
	nsongs, err := redisClient.LLen(songsKey).Result()
	if err != nil {
		ApiError(c, http.StatusInternalServerError, SongsGetError, fmt.Errorf("Unable to query for number of songs in playlist %s: %s", masterPlaylistId, err))
		return
	}
	var i int64
	songs := make([]models.Song, nsongs)
	for i = 0; i < nsongs; i += SONG_FETCH_BATCH_SIZE {
		songs_raw, err := redisClient.LRange(songsKey, i, i+SONG_FETCH_BATCH_SIZE).Result()
		if err != nil {
			ApiError(c, 500, SongsGetError, fmt.Errorf("Unable to fetch some songs: %s", err))
			return
		}
		for _, sr := range songs_raw {
			err := json.Unmarshal([]byte(sr), &songs[i])
			if err != nil {
				ApiError(c, 500, SongsGetError, fmt.Errorf("Unable to unmarshal song %s: %s", sr, err))
				return
			}
		}
	}
	log.Printf("Returning %d songs for playlist %s", len(songs), masterPlaylistId)

	c.JSON(http.StatusOK, songs)
}

func ApiV1EnqueueSong(c *gin.Context) {
	//add song to playlist
	requestedPlaylistId := c.Param("plid")

	//check if the requested plid is actually a contributor id and find the real playlist
	//TODO factor out this boilerplate of mapping a contributor ID to a master ID
	masterPlaylistId, err := redisClient.HGet("contributor_to_playlist", requestedPlaylistId).Result()
	if err != nil {
		ApiError(c, 500, SongEnqueueError, fmt.Errorf("Unable to lookup master playlist id from contributor id %s: %s", requestedPlaylistId, err))
		return
	}
	if masterPlaylistId == "" {
		masterPlaylistId = requestedPlaylistId
	}

	// decode the passed song to add to the list
	var s models.Song
	err = c.BindJSON(&s)
	if err != nil {
		ApiError(c, 422, err, err)
		return
	}

	songsKey := "songs:" + masterPlaylistId
	// TODO we should validate type and normalize the resource, but... fuck it. FIXME
	songjson, err := json.Marshal(s)
	if err != nil {
		ApiError(c, 422, err, err)
		return
	}
	i, err := redisClient.RPush(songsKey, string(songjson)).Result()
	if err != nil {
		ApiError(c, 500, SongEnqueueError, err)
	}

	log.Printf("Enqueued song %+v to playlist %s at %d", s, masterPlaylistId, i)
	c.JSON(http.StatusOK, gin.H{"playlist_length": i, "status": "enqueued"})
}

func ApiV1DequeueSong(c *gin.Context) {
	//delete song from playlist
	panic("TODO")
}
