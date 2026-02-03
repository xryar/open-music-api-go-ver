package web

type UserRegisterResponse struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
