package services

import (
	"context"
	"database/sql"
	"open-music-go/exception"
	"open-music-go/helper"
	"open-music-go/model/domain"
	web "open-music-go/model/web/playlist"
	playlistRepo "open-music-go/repositories/playlist"
	activityRepo "open-music-go/repositories/playlist_activity"
	songRepo "open-music-go/repositories/song"

	"github.com/go-playground/validator/v10"
)

type PlaylistServiceImpl struct {
	playlistRepository playlistRepo.PlaylistRepository
	songRepository     songRepo.SongRepository
	activityRepository activityRepo.PlaylistActivityRepository
	DB                 *sql.DB
	validate           *validator.Validate
}

func NewPlaylistService(
	playlistRepository playlistRepo.PlaylistRepository,
	songRepository songRepo.SongRepository,
	activityRepository activityRepo.PlaylistActivityRepository,
	db *sql.DB,
	validate *validator.Validate,
) *PlaylistServiceImpl {
	return &PlaylistServiceImpl{
		playlistRepository: playlistRepository,
		songRepository:     songRepository,
		activityRepository: activityRepository,
		DB:                 db,
		validate:           validate,
	}
}

func (ps *PlaylistServiceImpl) CreatePlaylist(ctx context.Context, request web.CreatePlaylistRequest) (web.PlaylistResponse, error) {
	err := ps.validate.Struct(request)
	helper.PanicIfError(err)

	userId, ok := ctx.Value("userId").(int)
	if !ok {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlist := domain.Playlist{
		Name:  request.Name,
		Owner: userId,
	}

	playlist = ps.playlistRepository.CreatePlaylist(ctx, tx, playlist)

	return helper.ToPlaylistResponse(playlist), nil
}

func (ps *PlaylistServiceImpl) AddSongToPlaylist(ctx context.Context, request web.PlaylistSongRequest) error {
	err := ps.validate.Struct(request)
	helper.PanicIfError(err)

	userId, ok := ctx.Value("userId").(int)
	if !ok {
		panic(exception.NewUnauthorizedError("unauthorized"))
	}

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlist, err := ps.playlistRepository.FindPlaylistById(ctx, tx, request.PlaylistId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if playlist.Owner != userId {
		panic(exception.NewUnauthorizedError("not owner"))
	}

	_, err = ps.songRepository.FindBySongId(ctx, tx, request.SongId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = ps.playlistRepository.AddSongToPlaylist(ctx, tx, request.PlaylistId, request.SongId)
	helper.PanicIfError(err)

	activity := domain.PlaylistActivity{
		PlaylistId: request.PlaylistId,
		SongId:     request.SongId,
		UserId:     userId,
		Action:     "ADD",
	}

	err = ps.activityRepository.Create(ctx, tx, activity)
	helper.PanicIfError(err)

	return nil
}

func (ps *PlaylistServiceImpl) DeleteSongInPlaylist(ctx context.Context, request web.PlaylistSongRequest) error {
	err := ps.validate.Struct(request)
	helper.PanicIfError(err)

	userId, ok := ctx.Value("userId").(int)
	if !ok {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlist, err := ps.playlistRepository.FindPlaylistById(ctx, tx, request.PlaylistId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if playlist.Owner != userId {
		panic(exception.NewUnauthorizedError("not owner"))
	}

	_, err = ps.songRepository.FindBySongId(ctx, tx, request.SongId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	ps.playlistRepository.DeleteSongInPlaylist(ctx, tx, request.PlaylistId, request.SongId)

	activity := domain.PlaylistActivity{
		PlaylistId: request.PlaylistId,
		SongId:     request.SongId,
		UserId:     userId,
		Action:     "DELETE",
	}

	err = ps.activityRepository.Create(ctx, tx, activity)

	return nil
}

func (ps *PlaylistServiceImpl) DeletePlaylist(ctx context.Context, id int) error {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		panic(exception.NewUnauthorizedError("unauthorized"))
	}

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlist, err := ps.playlistRepository.FindPlaylistById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if playlist.Owner != userId {
		panic(exception.NewUnauthorizedError("not owner"))
	}

	ps.playlistRepository.DeletePlaylist(ctx, tx, id)

	return nil
}

func (ps *PlaylistServiceImpl) FindPlaylistById(ctx context.Context, id int) (web.PlaylistResponse, error) {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlist, err := ps.playlistRepository.FindPlaylistById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPlaylistResponse(playlist), nil
}

func (ps *PlaylistServiceImpl) FindAllPlaylists(ctx context.Context) ([]web.PlaylistResponse, error) {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlists := ps.playlistRepository.FindAllPlaylists(ctx, tx)

	return helper.ToPlaylistResponses(playlists), nil
}

func (ps *PlaylistServiceImpl) FindPlaylistByOwner(ctx context.Context) ([]web.PlaylistResponse, error) {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		panic(exception.NewUnauthorizedError("unauthorized"))
	}

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlists := ps.playlistRepository.FindPlaylistByOwner(ctx, tx, userId)

	return helper.ToPlaylistResponses(playlists), nil
}
