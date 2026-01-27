package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SongController interface {
	CreateSong(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UpdateSong(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteSong(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindSongById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAllSong(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
