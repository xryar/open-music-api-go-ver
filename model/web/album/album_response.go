package web

import web "open-music-go/model/web/song"

type AlbumResponse struct {
	Id    int                `json:"id"`
	Name  string             `json:"name"`
	Year  int                `json:"year"`
	Songs []web.SongResponse `json:"songs,omitempty"`
}
