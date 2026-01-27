package services

import (
	"context"
	"open-music-go/model/domain"
	web "open-music-go/model/web/song"
)

type SongService interface {
	Create(ctx context.Context, request web.CreateSongRequest) domain.Song
	Update(ctx context.Context, request web.UpdateSongRequest) domain.Song
	Delete(ctx context.Context, id int)
	FindBySongId(ctx context.Context, id int) (domain.Song, error)
	FindAllSong(ctx context.Context) []web.SongResponse
}
