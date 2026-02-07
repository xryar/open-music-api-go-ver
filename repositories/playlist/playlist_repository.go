package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type PlaylistRepository interface {
	CreatePlaylist(ctx context.Context, tx *sql.Tx, playlist domain.Playlist) domain.Playlist
	DeletePlaylist(ctx context.Context, tx *sql.Tx, id int)
	FindPlaylistById(ctx context.Context, tx *sql.Tx, id int) []domain.Playlist
	FindAllPlaylists(ctx context.Context, tx *sql.Tx) (domain.Playlist, error)
}
