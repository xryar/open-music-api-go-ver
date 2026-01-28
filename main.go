package main

import (
	"fmt"
	"net/http"
	"open-music-go/app"
	controllers "open-music-go/controllers/album"
	"open-music-go/helper"
	repositories "open-music-go/repositories/album"
	services "open-music-go/services/album"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	albumRepository := repositories.NewAlbumRepository()
	albumService := services.NewAlbumService(albumRepository, db, validate)
	albumController := controllers.NewAlbumController(albumService)

	router := app.NewRouter(albumController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Starting web server at localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
