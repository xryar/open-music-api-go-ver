package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type PlaylistRepositoryImpl struct {
}

func NewPlaylistRepository() *PlaylistRepositoryImpl {
	return &PlaylistRepositoryImpl{}
}

func (pr *PlaylistRepositoryImpl) CreatePlaylist(ctx context.Context, tx *sql.Tx, playlist domain.Playlist) domain.Playlist {

}

func (pr *PlaylistRepositoryImpl) DeletePlaylist(ctx context.Context, tx *sql.Tx, id int) {

}

func (pr *PlaylistRepositoryImpl) FindPlaylistById(ctx context.Context, tx *sql.Tx, id int) []domain.Playlist {

}

func (pr *PlaylistRepositoryImpl) FindAllPlaylists(ctx context.Context, tx *sql.Tx) (domain.Playlist, error) {

}
