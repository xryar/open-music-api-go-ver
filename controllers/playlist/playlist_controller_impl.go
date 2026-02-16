package controllers

import (
	"net/http"
	services "open-music-go/services/playlist"

	"github.com/julienschmidt/httprouter"
)

type PlaylistControllerImpl struct {
	service services.PlaylistService
}

func NewPlaylistController(service services.PlaylistService) *PlaylistControllerImpl {
	return &PlaylistControllerImpl{service: service}
}

func (pc *PlaylistControllerImpl) CreatePlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (pc *PlaylistControllerImpl) AddSongToPlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (pc *PlaylistControllerImpl) UpdatePlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (pc *PlaylistControllerImpl) DeletePlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (pc *PlaylistControllerImpl) DeleteSongInPlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (pc *PlaylistControllerImpl) FindByPlaylistId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (pc *PlaylistControllerImpl) FindAllPlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
