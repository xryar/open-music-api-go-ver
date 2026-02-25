package services

import (
	"context"
	"database/sql"
	"open-music-go/exception"
	"open-music-go/helper"
	"open-music-go/model/domain"
	web "open-music-go/model/web/collab"
	playlistRepo "open-music-go/repositories/playlist"
	collabRepo "open-music-go/repositories/playlist_collab"
	userRepo "open-music-go/repositories/user"

	"github.com/go-playground/validator/v10"
)

type PlaylistCollabServiceImpl struct {
	CollabRepo   collabRepo.PlaylistCollabRepository
	PlaylistRepo playlistRepo.PlaylistRepository
	UserRepo     userRepo.UserRepository
	DB           *sql.DB
	Validate     *validator.Validate
}

func NewPlaylistCollabService(
	collabRepo collabRepo.PlaylistCollabRepository,
	playlistRepo playlistRepo.PlaylistRepository,
	userRepo userRepo.UserRepository,
	db *sql.DB,
	validate *validator.Validate,
) *PlaylistCollabServiceImpl {
	return &PlaylistCollabServiceImpl{
		CollabRepo:   collabRepo,
		PlaylistRepo: playlistRepo,
		UserRepo:     userRepo,
		DB:           db,
		Validate:     validate,
	}
}

func (pcs *PlaylistCollabServiceImpl) AddCollaborator(ctx context.Context, request web.CollabRequest) error {
	err := pcs.Validate.Struct(request)
	helper.PanicIfError(err)

	userId, ok := ctx.Value("userId").(int)
	if !ok {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	tx, err := pcs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlist, err := pcs.PlaylistRepo.FindPlaylistById(ctx, tx, request.PlaylistId)
	if err != nil {
		panic(exception.NewNotFoundError("playlist not found"))
	}

	if playlist.Owner != userId {
		panic(exception.NewUnauthorizedError("not owner"))
	}

	_, err = pcs.UserRepo.FindById(ctx, tx, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError("user not found"))
	}

	collab := domain.PlaylistCollab{
		PlaylistId: request.PlaylistId,
		UserId:     request.UserId,
	}

	exists, err := pcs.CollabRepo.IsCollab(ctx, tx, collab)
	helper.PanicIfError(err)

	if exists {
		panic(exception.NewBadRequestError("already collaborator"))
	}

	err = pcs.CollabRepo.Add(ctx, tx, collab)
	helper.PanicIfError(err)

	return nil
}

func (pcs *PlaylistCollabServiceImpl) RemoveCollaborator(ctx context.Context, request web.CollabRequest) error {
	err := pcs.Validate.Struct(request)
	helper.PanicIfError(err)

	userId, ok := ctx.Value("userId").(int)
	if !ok {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	tx, err := pcs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlist, err := pcs.PlaylistRepo.FindPlaylistById(ctx, tx, request.PlaylistId)
	if err != nil {
		panic(exception.NewNotFoundError("playlist not found"))
	}

	if playlist.Owner != userId {
		panic(exception.NewUnauthorizedError("not owner"))
	}

	collab := domain.PlaylistCollab{
		PlaylistId: request.PlaylistId,
		UserId:     request.UserId,
	}

	exists, err := pcs.CollabRepo.IsCollab(ctx, tx, collab)
	helper.PanicIfError(err)

	if !exists {
		panic(exception.NewNotFoundError("collaborator not found"))
	}

	err = pcs.CollabRepo.Remove(ctx, tx, collab)
	helper.PanicIfError(err)

	return nil
}

func (pcs *PlaylistCollabServiceImpl) GetCollaborators(ctx context.Context, playlistId int) ([]int, error) {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		panic(exception.NewUnauthorizedError("user not found"))
	}

	tx, err := pcs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	playlist, err := pcs.PlaylistRepo.FindPlaylistById(ctx, tx, playlistId)
	if err != nil {
		panic(exception.NewNotFoundError("playlist not found"))
	}

	collab := domain.PlaylistCollab{
		PlaylistId: playlistId,
		UserId:     userId,
	}

	isCollab, err := pcs.CollabRepo.IsCollab(ctx, tx, collab)
	helper.PanicIfError(err)

	if playlist.Owner != userId && !isCollab {
		panic(exception.NewUnauthorizedError("no access"))
	}

	collabs, err := pcs.CollabRepo.GetCollaborators(ctx, tx, playlistId)
	helper.PanicIfError(err)

	return helper.ToCollaboratorResponses(collabs), nil
}
