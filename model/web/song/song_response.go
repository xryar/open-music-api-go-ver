package web

type SongResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Performer string `json:"performer"`
	Genre     string `json:"genre"`
	Duration  int    `json:"duration"`
	AlbumId   int    `json:"albumId"`
}
