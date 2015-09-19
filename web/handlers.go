package web

import (
	"fmt"
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
	//TODO create a new playlist
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
