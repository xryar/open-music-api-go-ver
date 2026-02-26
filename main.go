package main

import (
	"fmt"
	"net/http"
	"open-music-go/app"
	albumController "open-music-go/controllers/album"
	playlistController "open-music-go/controllers/playlist"
	controllers "open-music-go/controllers/playlist_collab"
	songController "open-music-go/controllers/song"
	userController "open-music-go/controllers/user"
	"open-music-go/helper"
	albumRepository "open-music-go/repositories/album"
	playlistRepository "open-music-go/repositories/playlist"
	activityRepository "open-music-go/repositories/playlist_activity"
	collabRepository "open-music-go/repositories/playlist_collab"
	songRepository "open-music-go/repositories/song"
	userRepository "open-music-go/repositories/user"
	albumService "open-music-go/services/album"
	playlistService "open-music-go/services/playlist"
	collabService "open-music-go/services/playlist_collab"
	songService "open-music-go/services/song"
	userService "open-music-go/services/user"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	albumRepo := albumRepository.NewAlbumRepository()
	activityRepo := activityRepository.NewPlaylistActivityRepository()
	songRepo := songRepository.NewSongRepository()
	userRepo := userRepository.NewUserRepository()
	playlistRepo := playlistRepository.NewPlaylistRepository()
	collabRepo := collabRepository.NewPlaylistCollabRepository()
	albumService := albumService.NewAlbumService(albumRepo, db, validate)
	songService := songService.NewSongService(songRepo, db, validate)
	userService := userService.NewUserService(userRepo, db, validate)
	playlistService := playlistService.NewPlaylistService(playlistRepo, songRepo, activityRepo, db, validate)
	collabService := collabService.NewPlaylistCollabService(collabRepo, playlistRepo, userRepo, db, validate)
	albumController := albumController.NewAlbumController(albumService)
	songController := songController.NewSongController(songService)
	userController := userController.NewUserController(userService)
	playlistController := playlistController.NewPlaylistController(playlistService)
	collabController := controllers.NewPlaylistCollabController(collabService)

	router := app.NewRouter(
		albumController,
		songController,
		userController,
		playlistController,
		collabController,
	)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Starting web server at localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
