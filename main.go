package main

import (
	"fmt"
	"net/http"
	"open-music-go/app"
	albumController "open-music-go/controllers/album"
	songController "open-music-go/controllers/song"
	"open-music-go/helper"
	albumRepository "open-music-go/repositories/album"
	songRepository "open-music-go/repositories/song"
	albumService "open-music-go/services/album"
	songService "open-music-go/services/song"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	albumRepository := albumRepository.NewAlbumRepository()
	songRepository := songRepository.NewSongRepository()
	albumService := albumService.NewAlbumService(albumRepository, db, validate)
	songService := songService.NewSongService(songRepository, db, validate)
	albumController := albumController.NewAlbumController(albumService)
	songController := songController.NewSongController(songService)

	router := app.NewRouter(albumController, songController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Starting web server at localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
