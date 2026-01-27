package services

import (
	"context"
	web "open-music-go/model/web/song"
)

type SongService interface {
	CreateSong(ctx context.Context, request web.CreateSongRequest) web.SongResponse
	UpdateSong(ctx context.Context, request web.UpdateSongRequest) web.SongResponse
	DeleteSong(ctx context.Context, id int)
	FindBySongId(ctx context.Context, id int) web.SongResponse
	FindAllSong(ctx context.Context) []web.SongResponse
}
