package services

import (
	"context"
	web "open-music-go/model/web/album"
	repositories "open-music-go/repositories/album"
)

type AlbumServiceImpl struct {
	repository repositories.AlbumRepository
}

func NewAlbumService(repository repositories.AlbumRepository) *AlbumServiceImpl {
	return &AlbumServiceImpl{repository: repository}
}

func (as *AlbumServiceImpl) UpdateAlbum(ctx context.Context, request web.UpdateAlbumRequest) web.AlbumResponse {

}

func (as *AlbumServiceImpl) CreateAlbum(ctx context.Context, request web.CreateAlbumRequest) web.AlbumResponse {

}

func (as *AlbumServiceImpl) DeleteAlbum(ctx context.Context, id int) {

}

func (as *AlbumServiceImpl) FindByAlbumId(ctx context.Context, id int) web.AlbumResponse {

}

func (as *AlbumServiceImpl) FindAllAlbum(ctx context.Context) []web.AlbumResponse {

}
