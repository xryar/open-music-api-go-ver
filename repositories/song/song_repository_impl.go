package repositories

import (
	"context"
	"database/sql"
	"errors"
	"open-music-go/helper"
	"open-music-go/model/domain"
)

type SongRepositoryImpl struct {
}

func NewSongRepository() *SongRepositoryImpl {
	return &SongRepositoryImpl{}
}

func (sr *SongRepositoryImpl) CreateSong(ctx context.Context, tx *sql.Tx, song domain.Song) domain.Song {
	SQL := "INSERT INTO song(title, year, genre, performer, duration, album_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, song.Title, song.Year, song.Genre, song.Performer, song.Duration, song.AlbumId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	song.Id = int(id)

	return song
}

func (sr *SongRepositoryImpl) UpdateSong(ctx context.Context, tx *sql.Tx, song domain.Song) domain.Song {
	SQL := "UPDATE song SET title = ?, year = ?, genre = ?, performer = ?, duration = ?, album_id = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, song.Title, song.Year, song.Genre, song.Performer, song.Duration, song.AlbumId, song.Id)
	helper.PanicIfError(err)

	return song
}

func (sr *SongRepositoryImpl) DeleteSong(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM song WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (sr *SongRepositoryImpl) FindBySongId(ctx context.Context, tx *sql.Tx, id int) (domain.Song, error) {
	SQL := "SELECT id, title, year, genre, performer, duration, album_id FROM song WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	song := domain.Song{}
	if rows.Next() {
		err := rows.Scan(&song.Id, &song.Title, &song.Year, &song.Genre, &song.Performer, &song.Duration, &song.AlbumId)
		helper.PanicIfError(err)
		return song, nil
	} else {
		return song, errors.New("song not found")
	}
}

func (sr *SongRepositoryImpl) FindAllSong(ctx context.Context, tx *sql.Tx) []domain.Song {
	SQL := "SELECT id, title, year, genre, performer, duration, album_id FROM song"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var songs []domain.Song
	for rows.Next() {
		song := domain.Song{}
		err := rows.Scan(&song.Id, &song.Title, &song.Year, &song.Genre, &song.Performer, &song.Duration, &song.AlbumId)
		helper.PanicIfError(err)
		songs = append(songs, song)
	}

	return songs
}
