package models

type Playlist struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	//MasterHandle      string `json:"master_handle"`
	ContributorHandle string `json:"contributor_handle"`

	//Songs   []Song `json:"songs"` // this isnt right... TODO
	//OwnerId int64  `json:"owner_id"`
}

func (p Playlist) Public() bool {
	return true
}

func (p Playlist) Sharable() bool {
	return true
}
