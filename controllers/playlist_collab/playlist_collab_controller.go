package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PlaylistCollabController interface {
	AddCollaborator(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	RemoveCollaborator(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetAllCollaborators(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
