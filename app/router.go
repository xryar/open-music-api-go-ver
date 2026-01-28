package app

import (
	albumControllers "open-music-go/controllers/album"
	songControllers "open-music-go/controllers/song"
	"open-music-go/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	albumController albumControllers.AlbumController,
	songController songControllers.SongController,
) *httprouter.Router {
	router := httprouter.New()
	albumRouter(router, albumController)
	songRouter(router, songController)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func albumRouter(router *httprouter.Router, albumController albumControllers.AlbumController) {
	router.GET("/api/albums", albumController.FindAllAlbum)
	router.GET("/api/albums/:albumId", albumController.FindByAlbumId)
	router.POST("/api/albums", albumController.CreateAlbum)
	router.PUT("/api/albums/:albumId", albumController.UpdateAlbum)
	router.DELETE("/api/albums/:albumId", albumController.DeleteAlbum)
}

func songRouter(router *httprouter.Router, songController songControllers.SongController) {
	router.GET("/api/songs", songController.FindAllSong)
	router.GET("/api/songs/:songId", songController.FindSongById)
	router.POST("/api/songs", songController.CreateSong)
	router.PUT("/api/songs/:songId", songController.UpdateSong)
	router.DELETE("/api/songs/:songId", songController.DeleteSong)
}
