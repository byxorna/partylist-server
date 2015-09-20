package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/byxorna/partylist-server/models"
	"github.com/byxorna/partylist-server/util"
	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
)

var (
	DecodePlaylistError = errors.New("Unable to decode playlist")
	CreatePlaylistError = errors.New("Unable to create playlist")
	DeletePlaylistError = errors.New("Unable to delete playlist")
	GetPlaylistError    = errors.New("Unable to get playlist")
)

func ApiV1CreatePlaylist(c *gin.Context) {
	// create a new playlist
	var p models.Playlist

	// decode received JSON
	err := c.BindJSON(&p)
	if err != nil {
		ApiError(c, 422, DecodePlaylistError, err)
		return
	}

	p.Id = util.RandomId(30)
	p.ContributorKey = util.RandomId(30)

	log.Infof("Creating playlist %+v", p)

	_, err = redisClient.HMSet("playlist:"+p.Id,
		"id", p.Id,
		"name", p.Name,
		"owner", p.Owner,
		//	"master_handle", p.MasterHandle,
		"contributor_key", p.ContributorKey).Result()
	if err != nil {
		ApiError(c, 500, CreatePlaylistError, err)
		return
	}
	log.Infof("Created playlist at playlist:%s: %+v", p.Id, p)

	// track contributor handle mappings
	_, err = redisClient.HSet("contributor_to_playlist", p.ContributorKey, p.Id).Result()
	if err != nil {
		ApiError(c, 500, CreatePlaylistError, fmt.Errorf("Unable to map contributor handle to playlist: %s", p.Id, err))
		return
	}
	log.Infof("Mapped contributor key %s to playlist %s", p.ContributorKey, p.Id)

	//TODO is this necessary? just for tracking how many playlists we have
	_, err = redisClient.SAdd("playlists", p.Id).Result()
	if err != nil {
		ApiError(c, 500, CreatePlaylistError, fmt.Errorf("Unable to add playlist %s to playlist set: %s", p.Id, err))
		return
	}

	// create an empty songlist now
	redisClient.LPush("songs:"+p.Id, "empty")
	redisClient.LPop("songs:" + p.Id)

	c.JSON(http.StatusCreated, p)
}

func ApiV1GetPlaylist(c *gin.Context) {
	// get a playlist
	requestId := c.Param("plid")
	plid := ""
	// check if the plid is a contributor ID, substitute in the contributor handle for the ID
	res, _ := redisClient.HGet("contributor_to_playlist", requestId).Result()
	if res != "" {
		log.Infof("Contributor key %s maps to master playlist %s", requestId, res)
		plid = res
	} else {
		// no contributor key known, so this must be the master key
		plid = requestId
	}

	if !redisClient.Exists("playlist:" + plid).Val() {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	m, err := redisClient.HGetAllMap("playlist:" + plid).Result()
	if err != nil {
		ApiError(c, 500, GetPlaylistError, fmt.Errorf("Unable to fetch playlist %s: %s", plid, err))
		return
	}

	p := models.LoadPlaylistFromMap(m)
	log.Infof("Loaded playlist %+v", p)

	// replace master key with contributor key if necessary
	if p.ContributorKey == requestId {
		p.Id = requestId
	}

	c.JSON(http.StatusOK, p)
}

func ApiV1DeletePlaylist(c *gin.Context) {
	//TODO delete a playlist
}
