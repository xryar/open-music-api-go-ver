package web

import web "open-music-go/model/web/song"

type PlaylistResponse struct {
	Id    int                `json:"id"`
	Name  string             `json:"name"`
	Owner int                `json:"owner"`
	Songs []web.SongResponse `json:"songs,omitempty"`
}
