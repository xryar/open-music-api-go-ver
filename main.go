package main

import (
	"fmt"
	"net/http"
	"open-music-go/app"
	"open-music-go/helper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router, err := app.InitializeRouter()
	helper.PanicIfError(err)

	fmt.Println("Starting web server at localhost:3000")
	err = http.ListenAndServe(":3000", router)
	helper.PanicIfError(err)
}
