package repositories

import (
	"context"
	"database/sql"
	"open-music-go/helper"
	"open-music-go/model/domain"
)

type PlaylistRepositoryImpl struct {
}

func NewPlaylistRepository() *PlaylistRepositoryImpl {
	return &PlaylistRepositoryImpl{}
}

func (pr *PlaylistRepositoryImpl) CreatePlaylist(ctx context.Context, tx *sql.Tx, playlist domain.Playlist) domain.Playlist {
	SQL := "INSERT INTO playlists(name, owner) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, playlist.Name, playlist.Owner)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	playlist.Id = int(id)
	return playlist
}

func (pr *PlaylistRepositoryImpl) DeletePlaylist(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM playlist WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (pr *PlaylistRepositoryImpl) FindPlaylistById(ctx context.Context, tx *sql.Tx, id int) []domain.Playlist {

}

func (pr *PlaylistRepositoryImpl) FindAllPlaylists(ctx context.Context, tx *sql.Tx) (domain.Playlist, error) {

}
