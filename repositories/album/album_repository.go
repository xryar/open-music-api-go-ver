package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type AlbumRepository interface {
	CreateAlbum(ctx context.Context, tx *sql.Tx, album domain.Album) domain.Album
	UpdateAlbum(ctx context.Context, tx *sql.Tx, album domain.Album) domain.Album
	DeleteAlbum(ctx context.Context, tx *sql.Tx, album domain.Album)
	FindByAlbumId(ctx context.Context, tx *sql.Tx, id int) (domain.Album, error)
	FindAllAlbum(ctx context.Context, tx *sql.Tx) []domain.Album
}
