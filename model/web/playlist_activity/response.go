package web

import "time"

type PlaylistActivityResponse struct {
	PlaylistId int        `json:"playlistId"`
	Activities []Activity `json:"activities"`
}

type Activity struct {
	Username  string    `json:"username"`
	SongTitle string    `json:"songTitle"`
	Action    string    `json:"action"`
	Time      time.Time `json:"time"`
}
