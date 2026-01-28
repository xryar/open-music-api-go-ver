package app

import (
	controllers "open-music-go/controllers/album"
	"open-music-go/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(albumController controllers.AlbumController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/albums", albumController.FindAllAlbum)
	router.GET("/api/albums/:albumId", albumController.FindByAlbumId)
	router.POST("/api/albums", albumController.CreateAlbum)
	router.PUT("/api/albums/:albumId", albumController.UpdateAlbum)
	router.DELETE("/api/albums/:albumId", albumController.DeleteAlbum)

	router.PanicHandler = exception.ErrorHandler

	return router
}
