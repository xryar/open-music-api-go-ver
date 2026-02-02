package repositories

import (
	"context"
	"database/sql"
	"open-music-go/model/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
}
