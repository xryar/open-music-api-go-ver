package services

import (
	"context"
	web "open-music-go/model/web/album"
)

type AlbumService interface {
	CreateAlbum(ctx context.Context, request web.CreateAlbumRequest) web.AlbumResponse
	UpdateAlbum(ctx context.Context, request web.UpdateAlbumRequest) web.AlbumResponse
	DeleteAlbum(ctx context.Context, id int)
	FindByAlbumId(ctx context.Context, id int) web.AlbumResponse
	FindAllAlbum(ctx context.Context) []web.AlbumResponse
}
