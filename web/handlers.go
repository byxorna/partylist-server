package web

import (
	"encoding/json"
	"fmt"
	"github.com/byxorna/partylist-server/models"
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

	//TODO: get the logged in user id, set as owner id

	// decode received JSON
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		w.WriteHeader(422) // unprocessable entity
		ApiError(w, err)
		return
	}

	// Insert data to database
	stmt, err := db.Prepare("INSERT INTO playlists(name,owner_id) VALUES ($1,$2)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ApiError(w, err)
		return
	}
	_, err = stmt.Exec(p.Name, p.OwnerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ApiError(w, err)
		return
	}

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
