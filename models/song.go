package models

type Song struct {
	//TODO add ID, so we can dedupe songs by only storing their ID in the playlist, not the contents
	Type     SongType `json:"type" binding:"required"`     // youtube, soundcloud, etc
	Resource string   `json:"resource" binding:"required"` // resource identifier, i.e. URL, native id, whatever
}

type SongType int

const (
	SongTypeYoutube SongType = iota
	SongTypeSpotify
	SongTypeSoundcloud
)
