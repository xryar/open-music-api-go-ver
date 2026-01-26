package services

import (
	"context"
	"database/sql"
	"open-music-go/exception"
	"open-music-go/helper"
	"open-music-go/model/domain"
	web "open-music-go/model/web/album"
	repositories "open-music-go/repositories/album"

	"github.com/go-playground/validator/v10"
)

type AlbumServiceImpl struct {
	repository repositories.AlbumRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewAlbumService(repository repositories.AlbumRepository, db *sql.DB, validate *validator.Validate) *AlbumServiceImpl {
	return &AlbumServiceImpl{
		repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (as *AlbumServiceImpl) CreateAlbum(ctx context.Context, request web.CreateAlbumRequest) web.AlbumResponse {
	err := as.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	album := domain.Album{
		Name: request.Name,
		Year: request.Year,
	}

	album = as.repository.CreateAlbum(ctx, tx, album)

	return helper.ToAlbumResponse(album)
}

func (as *AlbumServiceImpl) UpdateAlbum(ctx context.Context, request web.UpdateAlbumRequest) web.AlbumResponse {
	err := as.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	album, err := as.repository.FindByAlbumId(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	album.Name = request.Name
	album.Year = request.Year

	album = as.repository.UpdateAlbum(ctx, tx, album)

	return helper.ToAlbumResponse(album)
}

func (as *AlbumServiceImpl) DeleteAlbum(ctx context.Context, id int) {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	album, err := as.repository.FindByAlbumId(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	as.repository.DeleteAlbum(ctx, tx, album)

}

func (as *AlbumServiceImpl) FindByAlbumId(ctx context.Context, id int) web.AlbumResponse {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	album, err := as.repository.FindByAlbumId(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAlbumResponse(album)
}

func (as *AlbumServiceImpl) FindAllAlbum(ctx context.Context) []web.AlbumResponse {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	albums := as.repository.FindAllAlbum(ctx, tx)

	return helper.ToAlbumResponses(albums)
}
