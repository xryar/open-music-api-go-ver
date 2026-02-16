package web

type CreatePlaylistRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

type SongToPlaylistRequest struct {
	SongId int `json:"songId" validate:"required,min=1"`
}

type PlaylistSongRequest struct {
	PlaylistId int
	SongId     int
}
