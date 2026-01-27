package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type SongRepository interface {
	CreateSong(ctx context.Context, tx *sql.Tx, song domain.Song) domain.Song
	UpdateSong(ctx context.Context, tx *sql.Tx, song domain.Song) domain.Song
	DeleteSong(ctx context.Context, tx *sql.Tx, id int)
	FindBySongId(ctx context.Context, tx *sql.Tx, id int) (domain.Song, error)
	FindAllSong(ctx context.Context, tx *sql.Tx) []domain.Song
}
