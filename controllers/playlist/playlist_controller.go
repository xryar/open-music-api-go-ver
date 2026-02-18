package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PlaylistController interface {
	CreatePlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	AddSongToPlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeletePlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteSongInPlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindByPlaylistId(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAllPlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindPlaylistByOwner(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
