package models

type Song struct {
	Id       int64
	Type     SongType // youtube, soundcloud, etc
	Resource string   // resource identifier, i.e. URL, native id, whatever
}

type SongType int

const (
	SongTypeYoutube SongType = iota
	SongTypeSpotify
	SongTypeSoundcloud
)
