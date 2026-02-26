package controllers

import (
	"net/http"
	"open-music-go/helper"
	web "open-music-go/model/web"
	req "open-music-go/model/web/collab"
	services "open-music-go/services/playlist_collab"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PlaylistCollabControllerImpl struct {
	service services.PlaylistCollabService
}

func NewPlaylistCollabController(service services.PlaylistCollabService) *PlaylistCollabControllerImpl {
	return &PlaylistCollabControllerImpl{service: service}
}

func (pcc *PlaylistCollabControllerImpl) AddCollaborator(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := req.CollabRequest{}
	helper.ReadFromRequestBody(r, &request)

	playlistId := p.ByName("playlistId")
	id, err := strconv.Atoi(playlistId)
	helper.PanicIfError(err)

	request.PlaylistId = id

	err = pcc.service.AddCollaborator(r.Context(), request)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    201,
		Status:  "OK",
		Message: "Collaborator Added",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (pcc *PlaylistCollabControllerImpl) RemoveCollaborator(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := req.CollabRequest{}
	helper.ReadFromRequestBody(r, &request)

	playlistId := p.ByName("playlistId")
	id, err := strconv.Atoi(playlistId)
	helper.PanicIfError(err)

	request.PlaylistId = id

	err = pcc.service.RemoveCollaborator(r.Context(), request)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Collaborator Removed",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (pcc *PlaylistCollabControllerImpl) GetAllCollaborators(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	playlistId := p.ByName("playlistId")
	id, err := strconv.Atoi(playlistId)
	helper.PanicIfError(err)

	collabs, err := pcc.service.GetCollaborators(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    201,
		Status:  "OK",
		Message: "Success Get Collaborators",
		Data:    collabs,
	}

	helper.WriteToResponseBody(w, webResponse)
}
