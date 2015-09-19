package models

type Playlist struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Songs   []Song `json:"songs"` // this isnt right... TODO
	OwnerId int64  `json:"owner_id"`
}

func (p Playlist) Public() bool {
	return true
}

func (p Playlist) Sharable() bool {
	return true
}
