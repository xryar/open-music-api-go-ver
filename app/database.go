package app

import (
	"database/sql"
	"open-music-go/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:acumalaka@tcp(localhost:3306)/open_music_go?parseTime=true&loc=Local")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
