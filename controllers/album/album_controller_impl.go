package controllers

import (
	"net/http"
	services "open-music-go/services/album"

	"github.com/julienschmidt/httprouter"
)

type AlbumControllerImpl struct {
	service services.AlbumService
}

func NewAlbumService(service services.AlbumService) *AlbumControllerImpl {
	return &AlbumControllerImpl{service: service}
}

func (ac *AlbumControllerImpl) CreateAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ac *AlbumControllerImpl) UpdateAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ac *AlbumControllerImpl) DeleteAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ac *AlbumControllerImpl) FindByAlbumId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ac *AlbumControllerImpl) FindAllAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
