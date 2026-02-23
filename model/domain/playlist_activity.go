package domain

import "time"

type PlaylistActivity struct {
	Id         int
	PlaylistId int
	SongId     int
	UserId     int
	Action     string
	Time       time.Time
}

type PlaylistActivityJoin struct {
	Username  string
	SongTitle string
	Action    string
	Time      time.Time
}
