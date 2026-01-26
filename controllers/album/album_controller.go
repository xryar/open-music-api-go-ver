package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AlbumController interface {
	CreateAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UpdateAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindByAlbumId(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAllAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
