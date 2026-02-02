package repositories

import (
	"context"
	"database/sql"
	"open-music-go/helper"
	"open-music-go/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (us *UserRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(fullname, username, password) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Fullname, user.Username, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}
