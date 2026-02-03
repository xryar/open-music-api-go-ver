package controllers

import (
	"net/http"
	"open-music-go/helper"
	"open-music-go/model/web"
	req "open-music-go/model/web/user"
	service "open-music-go/services/user"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{service: service}
}

func (uc *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userRequest := req.UserRegisterRequest{}
	helper.ReadFromRequestBody(r, &userRequest)

	response := uc.service.Register(r.Context(), userRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (uc *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userRequest := req.UserLoginRequest{}
	helper.ReadFromRequestBody(r, &userRequest)

	response := uc.service.Login(r.Context(), userRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}
