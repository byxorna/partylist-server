package models

type Playlist struct {
	Id    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Owner string `json:"owner" binding:"required"`
	//MasterHandle      string `json:"master_handle"`
	ContributorKey string `json:"contributor_key"`

	//Songs   []Song `json:"songs"` // this isnt right... TODO
	//OwnerId int64  `json:"owner_id"`
}

func (p Playlist) Public() bool {
	return true
}

func (p Playlist) Sharable() bool {
	return true
}

//TODO there has GOT to be a better way to do this.
// Load a playlist model from the hash handed back from redis HGetAllMap
func LoadPlaylistFromMap(m map[string]string) Playlist {
	p := Playlist{}
	p.Id = m["id"]
	p.Name = m["name"]
	p.Owner = m["owner"]
	p.ContributorKey = m["contributor_key"]
	return p

}
