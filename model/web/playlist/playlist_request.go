package web

type CreatePlaylistRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

type AddSongToPlaylistRequest struct {
	PlaylistId int `json:"playlistId" validate:"required,min=1"`
	SongId     int `json:"songId" validate:"required,min=1"`
}
