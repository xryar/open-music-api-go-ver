package repositories

import (
	"context"
	"database/sql"
	"errors"
	"open-music-go/exception"
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

func (pr *PlaylistRepositoryImpl) AddSongToPlaylist(ctx context.Context, tx *sql.Tx, playlistId, songId int) error {
	SQL := "INSERT INTO playlist_song(playlist_id, song_id) VALUES (?, ?)"
	_, err := tx.ExecContext(ctx, SQL, playlistId, songId)
	helper.PanicIfError(err)

	return nil
}

func (pr *PlaylistRepositoryImpl) DeleteSongInPlaylist(ctx context.Context, tx *sql.Tx, playlistId, songId int) {
	SQL := "DELETE FROM playlist_songs WHERE playlist_id = ? AND song_id = ?"
	result, err := tx.ExecContext(ctx, SQL, playlistId, songId)
	helper.PanicIfError(err)

	rows, err := result.RowsAffected()
	helper.PanicIfError(err)

	if rows == 0 {
		panic(exception.NewNotFoundError("song not found in playlist"))
	}
}

func (pr *PlaylistRepositoryImpl) FindPlaylistByOwner(ctx context.Context, tx *sql.Tx, userId int) []domain.Playlist {
	SQL := "SELECT id, name, owner FROM playlists WHERE owner = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var playlists []domain.Playlist
	for rows.Next() {
		playlist := domain.Playlist{}
		err := rows.Scan(&playlist.Id, &playlist.Name, &playlist.Owner)
		helper.PanicIfError(err)
		playlists = append(playlists, playlist)
	}

	return playlists
}

func (pr *PlaylistRepositoryImpl) DeletePlaylist(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM playlist WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (pr *PlaylistRepositoryImpl) FindPlaylistById(ctx context.Context, tx *sql.Tx, id int) (domain.Playlist, error) {
	SQL := `
	SELECT
		p.id, p.name, p.owner,
		s.id, s.title, s.year, s.genre, s.performer, s.duration
	FROM playlists p
	LEFT JOIN playlist_songs ps on ps.playlist_id = p.id
	LEFT JOIN song s ON s.id = ps.song_id
	WHERE p.id = ?
	`
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	playlist := domain.Playlist{
		Songs: []domain.Song{},
	}
	found := false

	for rows.Next() {
		found = true

		var songId sql.NullInt64
		var title sql.NullString
		var year sql.NullInt64
		var performer sql.NullString
		var genre sql.NullString
		var duration sql.NullInt64

		err := rows.Scan(
			&playlist.Id,
			&playlist.Name,
			&playlist.Owner,
			&songId,
			&title,
			&year,
			&performer,
			&genre,
			&duration,
		)
		helper.PanicIfError(err)

		if !songId.Valid {
			continue
		}

		song := domain.Song{
			Id:        int(songId.Int64),
			Title:     title.String,
			Year:      int(year.Int64),
			Performer: performer.String,
			Genre:     genre.String,
			Duration:  int(duration.Int64),
		}
		playlist.Songs = append(playlist.Songs, song)
	}

	if !found {
		return playlist, errors.New("playlist not found")
	}

	return playlist, nil
}

func (pr *PlaylistRepositoryImpl) FindAllPlaylists(ctx context.Context, tx *sql.Tx) []domain.Playlist {
	SQL := "SELECT id, name, owner FROM playlists"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var playlists []domain.Playlist
	for rows.Next() {
		playlist := domain.Playlist{}
		err := rows.Scan(&playlist.Id, &playlist.Name, &playlist.Owner)
		helper.PanicIfError(err)
		playlists = append(playlists, playlist)
	}

	return playlists
}
