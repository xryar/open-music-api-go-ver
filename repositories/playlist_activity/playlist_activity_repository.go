package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type PlaylistActivityRepository interface {
	Create(ctx context.Context, tx *sql.Tx, activity domain.PlaylistActivity) error
	FindPlaylistById(ctx context.Context, tx *sql.Tx, playlistId int) ([]domain.PlaylistActivityJoin, error)
}
