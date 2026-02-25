package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type PlaylistCollabRepository interface {
	Add(ctx context.Context, tx *sql.Tx, collaborator domain.PlaylistCollab) error
	Remove(ctx context.Context, tx *sql.Tx, collaborator domain.PlaylistCollab) error
	IsCollab(ctx context.Context, tx *sql.Tx, collaborator domain.PlaylistCollab) (bool, error)
	GetCollaborators(ctx context.Context, tx *sql.Tx, playlistId int) ([]int, error)
}
