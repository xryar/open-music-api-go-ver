package services

import (
	"context"
	web "open-music-go/model/web/collab"
)

type PlaylistCollabService interface {
	AddCollaborator(ctx context.Context, request web.CollabRequest) error
	RemoveCollaborator(ctx context.Context, request web.CollabRequest) error
	GetCollaborators(ctx context.Context, playlistId int) ([]web.CollabResponse, error)
}
