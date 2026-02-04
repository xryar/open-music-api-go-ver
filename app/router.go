package app

import (
	albumControllers "open-music-go/controllers/album"
	songControllers "open-music-go/controllers/song"
	userControllers "open-music-go/controllers/user"
	"open-music-go/exception"
	"open-music-go/middlewares"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	albumController albumControllers.AlbumController,
	songController songControllers.SongController,
	userController userControllers.UserController,
) *httprouter.Router {
	router := httprouter.New()
	albumRouter(router, albumController)
	songRouter(router, songController)
	userRouter(router, userController)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func albumRouter(router *httprouter.Router, albumController albumControllers.AlbumController) {
	router.GET("/api/albums", middlewares.AuthMiddleware(albumController.FindAllAlbum))
	router.GET("/api/albums/:albumId", middlewares.AuthMiddleware(albumController.FindByAlbumId))
	router.POST("/api/albums", middlewares.AuthMiddleware(albumController.CreateAlbum))
	router.PUT("/api/albums/:albumId", middlewares.AuthMiddleware(albumController.UpdateAlbum))
	router.DELETE("/api/albums/:albumId", middlewares.AuthMiddleware(albumController.DeleteAlbum))
}

func songRouter(router *httprouter.Router, songController songControllers.SongController) {
	router.GET("/api/songs", middlewares.AuthMiddleware(songController.FindAllSong))
	router.GET("/api/songs/:songId", middlewares.AuthMiddleware(songController.FindSongById))
	router.POST("/api/songs", middlewares.AuthMiddleware(songController.CreateSong))
	router.PUT("/api/songs/:songId", middlewares.AuthMiddleware(songController.UpdateSong))
	router.DELETE("/api/songs/:songId", middlewares.AuthMiddleware(songController.DeleteSong))
}

func userRouter(router *httprouter.Router, userController userControllers.UserController) {
	router.POST("/api/user/register", userController.Register)
	router.POST("/api/user/login", userController.Login)
}
