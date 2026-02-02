package web

type UserRegisterResponse struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
