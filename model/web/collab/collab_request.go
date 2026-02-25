package web

type CollabRequest struct {
	PlaylistId int `json:"playlistId" validate:"required"`
	UserId     int `json:"userId" validate:"required"`
}
