package service

import (
	"context"
	"database/sql"
	"open-music-go/helper"
	"open-music-go/model/domain"
	web "open-music-go/model/web/user"
	repositories "open-music-go/repositories/user"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository repositories.UserRepository
	DB         *sql.DB
	validate   *validator.Validate
}

func NewUserUservice(repository repositories.UserRepository, db *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repository,
		DB:         db,
		validate:   validate,
	}
}

func (us *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) web.UserRegisterResponse {
	err := us.validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	user := domain.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Password: string(hashedPassword),
	}

	user = us.repository.CreateUser(ctx, tx, user)

	return helper.ToUserRegisterResponse(user)
}

func (us *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse {
	err := us.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollack(tx)

	user, err := us.repository.FindByUsername(ctx, tx, request.Username)
	helper.PanicIfError(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	helper.PanicIfError(err)

	token, err := helper.GenerateToken(user.Id, user.Username)
	helper.PanicIfError(err)

	return helper.ToUserLoginResponse(token)
}
