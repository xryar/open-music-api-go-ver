package app

import (
	albumControllers "open-music-go/controllers/album"
	playlistControllers "open-music-go/controllers/playlist"
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
	playlistController playlistControllers.PlaylistController,
) *httprouter.Router {
	router := httprouter.New()
	albumRouter(router, albumController)
	songRouter(router, songController)
	userRouter(router, userController)
	playlistRouter(router, playlistController)

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

func userRouter(router *httprouter.Router, userController userControllers.UserController) {
	router.POST("/api/user/register", userController.Register)
	router.POST("/api/user/login", userController.Login)
}

func playlistRouter(router *httprouter.Router, playlistController playlistControllers.PlaylistController) {
	router.POST("/api/playlists", middlewares.AuthMiddleware(playlistController.CreatePlaylist))
	router.GET("/api/playlists/me", middlewares.AuthMiddleware(playlistController.FindPlaylistByOwner))
	router.GET("/api/playlists", playlistController.FindAllPlaylist)
	router.DELETE("/api/playlists/:id", middlewares.AuthMiddleware(playlistController.DeletePlaylist))
	router.POST("/api/playlists/:id/songs", middlewares.AuthMiddleware(playlistController.AddSongToPlaylist))
	router.GET("/api/playlists/:id/songs", middlewares.AuthMiddleware(playlistController.FindByPlaylistId))
	router.DELETE("/api/playlists/:id/songs", middlewares.AuthMiddleware(playlistController.DeleteSongInPlaylist))
}
