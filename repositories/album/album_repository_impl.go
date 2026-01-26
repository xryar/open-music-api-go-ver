package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type AlbumRepositoryImpl struct {
}

func NewAlbumRepository() AlbumRepositoryImpl {
	return AlbumRepositoryImpl{}
}

func (ar *AlbumRepositoryImpl) CreateAlbum(ctx context.Context, tx *sql.Tx, album domain.Album) domain.Album {
}

func (ar *AlbumRepositoryImpl) UpdateAlbum(ctx context.Context, tx *sql.Tx, album domain.Album) domain.Album {
}

func (ar *AlbumRepositoryImpl) DeleteAlbum(ctx context.Context, tx *sql.Tx, album domain.Album) {}

func (ar *AlbumRepositoryImpl) FindByAlbumId(ctx context.Context, tx *sql.Tx, id int) (domain.Album, error) {
}

func (ar *AlbumRepositoryImpl) FindAllAlbum(ctx context.Context, tx *sql.Tx) []domain.Album {}
