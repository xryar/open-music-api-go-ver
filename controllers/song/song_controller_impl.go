package controllers

import (
	"net/http"
	"open-music-go/helper"
	"open-music-go/model/web"
	req "open-music-go/model/web/song"
	services "open-music-go/services/song"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type SongControllerImpl struct {
	service services.SongService
}

func NewSongController(service services.SongService) *SongControllerImpl {
	return &SongControllerImpl{service: service}
}

func (sc *SongControllerImpl) CreateSong(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songCreateRequest := req.CreateSongRequest{}
	helper.ReadFromRequestBody(r, &songCreateRequest)

	songResponse := sc.service.CreateSong(r.Context(), songCreateRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   songResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (sc *SongControllerImpl) UpdateSong(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songUpdateRequest := req.UpdateSongRequest{}
	helper.ReadFromRequestBody(r, &songUpdateRequest)

	songId := p.ByName("songId")
	id, err := strconv.Atoi(songId)
	helper.PanicIfError(err)

	songUpdateRequest.Id = id

	songResponse := sc.service.UpdateSong(r.Context(), songUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   songResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (sc *SongControllerImpl) DeleteSong(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songId := p.ByName("songId")
	id, err := strconv.Atoi(songId)
	helper.PanicIfError(err)

	sc.service.DeleteSong(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (sc *SongControllerImpl) FindSongById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songId := p.ByName("songId")
	id, err := strconv.Atoi(songId)
	helper.PanicIfError(err)

	songResponse := sc.service.FindBySongId(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   songResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (sc *SongControllerImpl) FindAllSong(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	songResponses := sc.service.FindAllSong(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   songResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
