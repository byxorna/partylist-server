package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/byxorna/partylist-server/models"
	"github.com/byxorna/partylist-server/util"
	log "github.com/golang/glog"
	"github.com/gorilla/mux"
)

var (
	DecodePlaylistError = errors.New("Unable to decode playlist")
	CreatePlaylistError = errors.New("Unable to create playlist")
	DeletePlaylistError = errors.New("Unable to delete playlist")
	GetPlaylistError    = errors.New("Unable to get playlist")
)

func ApiV1CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	// create a new playlist
	var p models.Playlist

	// decode received JSON
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		ApiError(w, 422, DecodePlaylistError, err)
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
		ApiError(w, 500, CreatePlaylistError, err)
		return
	}
	log.Infof("Created playlist at playlist:%s: %+v", p.Id, p)

	// track contributor handle mappings
	_, err = redisClient.HSet("contributor_to_playlist", p.ContributorKey, p.Id).Result()
	if err != nil {
		ApiError(w, 500, CreatePlaylistError, fmt.Errorf("Unable to map contributor handle to playlist: %s", p.Id, err))
		return
	}
	log.Infof("Mapped contributor key %s to playlist %s", p.ContributorKey, p.Id)

	//TODO is this necessary? just for tracking how many playlists we have
	_, err = redisClient.SAdd("playlists", p.Id).Result()
	if err != nil {
		ApiError(w, 500, CreatePlaylistError, fmt.Errorf("Unable to add playlist %s to playlist set: %s", p.Id, err))
		return
	}

	// create an empty songlist now
	redisClient.LPush("songs:"+p.Id, "empty")
	redisClient.LPop("songs:" + p.Id)

	//TODO too much copypasta
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}

}

func ApiV1GetPlaylist(w http.ResponseWriter, r *http.Request) {
	// get a playlist
	vars := mux.Vars(r)
	requestId := vars["plid"]
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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	m, err := redisClient.HGetAllMap("playlist:" + plid).Result()
	if err != nil {
		ApiError(w, 500, GetPlaylistError, fmt.Errorf("Unable to fetch playlist %s: %s", plid, err))
		return
	}

	p := models.LoadPlaylistFromMap(m)
	log.Infof("Loaded playlist %+v", p)

	// replace master key with contributor key if necessary
	if p.ContributorKey == requestId {
		p.Id = requestId
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		ApiError(w, 500, GetPlaylistError, fmt.Errorf("Unable to encode playlist %s: %s", p.Id, err))
		return
	}
}

func ApiV1DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	//TODO delete a playlist
}
