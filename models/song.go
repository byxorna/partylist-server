package models

type Song struct {
	Type     SongType `json:"type" binding:"required"`     // youtube, soundcloud, etc
	Resource string   `json:"resource" binding:"required"` // resource identifier, i.e. URL, native id, whatever
}

type SongType int

const (
	SongTypeYoutube SongType = iota
	SongTypeSpotify
	SongTypeSoundcloud
)

//TODO this sucks and i hate it. figure out a better way of reviving a model from redis
func LoadSongFromMap(m map[string]string) Song {
	s := Song{}
	s.Resource = m["resource"]
	return s
}
