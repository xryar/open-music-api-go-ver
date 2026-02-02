package repositories

import (
	"context"
	"database/sql"
	"errors"
	"open-music-go/helper"
	"open-music-go/model/domain"
)

type AlbumRepositoryImpl struct {
}

func NewAlbumRepository() *AlbumRepositoryImpl {
	return &AlbumRepositoryImpl{}
}

func (ar *AlbumRepositoryImpl) CreateAlbum(ctx context.Context, tx *sql.Tx, album domain.Album) domain.Album {
	SQL := "INSERT INTO album(name, year) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, SQL, album.Name, album.Year)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	album.Id = int(id)

	return album
}

func (ar *AlbumRepositoryImpl) UpdateAlbum(ctx context.Context, tx *sql.Tx, album domain.Album) domain.Album {
	SQL := "UPDATE album SET name = ?, year = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, album.Name, album.Year, album.Id)
	helper.PanicIfError(err)

	return album
}

func (ar *AlbumRepositoryImpl) DeleteAlbum(ctx context.Context, tx *sql.Tx, album domain.Album) {
	SQL := "DELETE FROM album WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, album.Id)
	helper.PanicIfError(err)
}

func (ar *AlbumRepositoryImpl) FindByAlbumId(ctx context.Context, tx *sql.Tx, id int) (domain.Album, error) {
	SQL := "SELECT a.id, a.name, a.year, s.id, s.title, s.year, s.performer, s.genre, s.duration FROM album AS a LEFT JOIN song s ON s.album_id = a.id WHERE a.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	album := domain.Album{
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
			&album.Id,
			&album.Name,
			&album.Year,
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

		album.Songs = append(album.Songs, song)
	}

	if !found {
		return album, errors.New("album not found")
	}

	return album, nil
}

func (ar *AlbumRepositoryImpl) FindAllAlbum(ctx context.Context, tx *sql.Tx) []domain.Album {
	SQL := "SELECT id, name, year FROM album"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var albums []domain.Album
	for rows.Next() {
		album := domain.Album{}
		err := rows.Scan(&album.Id, &album.Name, &album.Year)
		helper.PanicIfError(err)
		albums = append(albums, album)
	}

	return albums
}
