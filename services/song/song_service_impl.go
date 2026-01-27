package services

import (
	"context"
	"database/sql"
	"open-music-go/exception"
	"open-music-go/helper"
	"open-music-go/model/domain"
	web "open-music-go/model/web/song"
	repositories "open-music-go/repositories/song"

	"github.com/go-playground/validator/v10"
)

type SongServiceImpl struct {
	repository repositories.SongRepository
	DB         *sql.DB
	validate   *validator.Validate
}

func NewSongService(repository repositories.SongRepository, db *sql.DB, validate *validator.Validate) *SongServiceImpl {
	return &SongServiceImpl{
		repository: repository,
		DB:         db,
		validate:   validate,
	}
}

func (ss *SongServiceImpl) CreateSong(ctx context.Context, request web.CreateSongRequest) web.SongResponse {
	err := ss.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ss.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	song := domain.Song{
		Title:     request.Title,
		Year:      request.Year,
		Genre:     request.Genre,
		Performer: request.Performer,
		Duration:  request.Duration,
		AlbumId:   request.AlbumId,
	}

	song = ss.repository.CreateSong(ctx, tx, song)

	return helper.ToSongResponse(song)
}

func (ss *SongServiceImpl) UpdateSong(ctx context.Context, request web.UpdateSongRequest) web.SongResponse {
	err := ss.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ss.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	song, err := ss.repository.FindBySongId(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	song = domain.Song{
		Title:     request.Title,
		Year:      request.Year,
		Performer: request.Performer,
		Genre:     request.Genre,
		Duration:  request.Duration,
		AlbumId:   request.AlbumId,
	}

	song = ss.repository.UpdateSong(ctx, tx, song)

	return helper.ToSongResponse(song)
}

func (ss *SongServiceImpl) DeleteSong(ctx context.Context, id int) {
	tx, err := ss.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	song, err := ss.repository.FindBySongId(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	ss.repository.DeleteSong(ctx, tx, song.Duration)
}

func (ss *SongServiceImpl) FindBySongId(ctx context.Context, id int) web.SongResponse {
	tx, err := ss.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	song, err := ss.repository.FindBySongId(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSongResponse(song)
}

func (ss *SongServiceImpl) FindAllSong(ctx context.Context) []web.SongResponse {
	tx, err := ss.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	songs := ss.repository.FindAllSong(ctx, tx)

	return helper.ToSongResponses(songs)
}
