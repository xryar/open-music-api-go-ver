package controllers

import (
	"net/http"
	"open-music-go/helper"
	"open-music-go/model/web"
	req "open-music-go/model/web/album"
	services "open-music-go/services/album"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AlbumControllerImpl struct {
	service services.AlbumService
}

func NewAlbumController(service services.AlbumService) *AlbumControllerImpl {
	return &AlbumControllerImpl{service: service}
}

func (ac *AlbumControllerImpl) CreateAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumCreateRequest := req.CreateAlbumRequest{}
	helper.ReadFromRequestBody(r, &albumCreateRequest)

	albumResponse := ac.service.CreateAlbum(r.Context(), albumCreateRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   albumResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (ac *AlbumControllerImpl) UpdateAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumUpdateRequest := req.UpdateAlbumRequest{}
	helper.ReadFromRequestBody(r, &albumUpdateRequest)

	albumId := p.ByName("albumId")
	id, err := strconv.Atoi(albumId)
	helper.PanicIfError(err)

	albumUpdateRequest.Id = id

	albumResponse := ac.service.UpdateAlbum(r.Context(), albumUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   albumResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (ac *AlbumControllerImpl) DeleteAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumId := p.ByName("albumId")
	id, err := strconv.Atoi(albumId)
	helper.PanicIfError(err)

	ac.service.DeleteAlbum(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (ac *AlbumControllerImpl) FindByAlbumId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumId := p.ByName("albumId")
	id, err := strconv.Atoi(albumId)
	helper.PanicIfError(err)

	albumResponse := ac.service.FindByAlbumId(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   albumResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (ac *AlbumControllerImpl) FindAllAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	albumResponses := ac.service.FindAllAlbum(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   albumResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
