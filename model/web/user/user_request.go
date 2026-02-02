package web

type UserRegisterRequest struct {
	Fullname string `json:"fullname" validate:"required,min=4,max255"`
	Username string `json:"username" validate:"required,min=4,max=255"`
	Password string `json:"password" validate:"required,min=6,max=25"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,min=4,max=255"`
	Password string `json:"password" validate:"required,min=6,max=25"`
}
