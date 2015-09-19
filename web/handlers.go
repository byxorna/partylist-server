package web

import (
	"encoding/json"
	"fmt"
	"github.com/byxorna/partylist-server/models"
	"github.com/byxorna/partylist-server/util"
	log "github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ApiV1Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API v1 Index")
}

func ApiV1CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	// create a new playlist
	var p models.Playlist

	// decode received JSON
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		w.WriteHeader(422) // unprocessable entity
		ApiError(w, err)
		return
	}

	p.Id = util.RandomId(30)
	//p.MasterHandle = util.RandomId(30)
	p.ContributorHandle = util.RandomId(30)

	log.Infof("Creating playlist %+v", p)

	_, err = redisClient.HMSet("playlist:"+p.Id,
		"name", p.Name,
		"owner", p.Owner,
		//	"master_handle", p.MasterHandle,
		"contributor_handle", p.ContributorHandle).Result()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("Unable to create playlist %s: %s", p.Id, err)
		ApiError(w, err)
		return
	}
	log.Infof("Created playlist at playlist:%s: %+v", p.Id, p)

	// track contributor handle mappings
	_, err = redisClient.HSet("contributor_to_playlist", p.ContributorHandle, p.Id).Result()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("Unable to map contributor handle to playlist: %s", p.Id, err)
		ApiError(w, err)
		return
	}
	log.Infof("Mapped contributor handle %s to playlist %s", p.ContributorHandle, p.Id)

	//TODO is this necessary? just for tracking how many playlists we have
	_, err = redisClient.SAdd("playlists", p.Id).Result()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("Unable to add playlist %s to playlist set: %s", p.Id, err)
		ApiError(w, err)
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
	//TODO show a playlist
	vars := mux.Vars(r)
	id := vars["plid"]
	fmt.Fprintln(w, "Playlist show:", id)
}

func ApiV1DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	//TODO delete a playlist
}

func ApiV1AddSong(w http.ResponseWriter, r *http.Request) {
	//TODO add song to playlist
}

func ApiV1DeleteSong(w http.ResponseWriter, r *http.Request) {
	//TODO delete song from playlist
}
