package service

import (
	"context"
	web "open-music-go/model/web/user"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) web.UserRegisterResponse
	Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
}
