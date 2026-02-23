package controllers

import (
	"net/http"
	"open-music-go/helper"
	"open-music-go/model/web"
	req "open-music-go/model/web/playlist"
	services "open-music-go/services/playlist"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PlaylistControllerImpl struct {
	service services.PlaylistService
}

func NewPlaylistController(service services.PlaylistService) *PlaylistControllerImpl {
	return &PlaylistControllerImpl{service: service}
}

func (pc *PlaylistControllerImpl) CreatePlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	playlistRequest := req.CreatePlaylistRequest{}
	helper.ReadFromRequestBody(r, &playlistRequest)

	playlist, err := pc.service.CreatePlaylist(r.Context(), playlistRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    201,
		Status:  "OK",
		Message: "Success Create Playlist",
		Data:    playlist,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (pc *PlaylistControllerImpl) AddSongToPlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bodyRequest := req.SongToPlaylistRequest{}
	helper.ReadFromRequestBody(r, &bodyRequest)

	playlistId := p.ByName("playlistId")
	id, err := strconv.Atoi(playlistId)
	helper.PanicIfError(err)

	request := req.PlaylistSongRequest{
		PlaylistId: id,
		SongId:     bodyRequest.SongId,
	}

	err = pc.service.AddSongToPlaylist(r.Context(), request)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    201,
		Status:  "OK",
		Message: "Success Add Song To Playlist",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (pc *PlaylistControllerImpl) DeletePlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	playlistId := p.ByName("playlistId")
	id, err := strconv.Atoi(playlistId)
	helper.PanicIfError(err)

	err = pc.service.DeletePlaylist(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Delete Playlist",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (pc *PlaylistControllerImpl) DeleteSongInPlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bodyRequest := req.SongToPlaylistRequest{}
	helper.ReadFromRequestBody(r, &bodyRequest)

	playlistId := p.ByName("playlistId")
	id, err := strconv.Atoi(playlistId)
	helper.PanicIfError(err)

	request := req.PlaylistSongRequest{
		PlaylistId: id,
		SongId:     bodyRequest.SongId,
	}

	err = pc.service.DeleteSongInPlaylist(r.Context(), request)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Song removed from playlist",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (pc *PlaylistControllerImpl) FindByPlaylistId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	playlistId := p.ByName("playlistId")
	id, err := strconv.Atoi(playlistId)
	helper.PanicIfError(err)

	playlist, err := pc.service.FindPlaylistById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Get Playlist",
		Data:    playlist,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (pc *PlaylistControllerImpl) FindPlaylistByOwner(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	playlists, err := pc.service.FindPlaylistByOwner(r.Context())
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Get Playlist by Owner",
		Data:    playlists,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (pc *PlaylistControllerImpl) FindAllPlaylist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	playlists, err := pc.service.FindAllPlaylists(r.Context())
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Get All Playlists",
		Data:    playlists,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (pc *PlaylistControllerImpl) GetPlaylistActivities(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	playlistId := p.ByName("playlistId")
	id, err := strconv.Atoi(playlistId)
	helper.PanicIfError(err)

	activities, err := pc.service.GetPlaylistActivities(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Get Playlist Activities",
		Data:    activities,
	}

	helper.WriteToResponseBody(w, webResponse)
}
