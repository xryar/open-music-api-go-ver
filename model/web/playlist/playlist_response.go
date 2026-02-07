package web

import web "open-music-go/model/web/song"

type PlaylistResponse struct {
	Id       int                `json:"id"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Songs    []web.SongResponse `json:"songs,omitempty"`
}
