package helper

import (
	"open-music-go/model/domain"
	web "open-music-go/model/web/album"
	web2 "open-music-go/model/web/song"
	web3 "open-music-go/model/web/user"
)

func ToAlbumResponse(album domain.Album) web.AlbumResponse {
	songs := make([]web2.SongResponse, 0)

	for _, song := range album.Songs {
		songs = append(songs, web2.SongResponse{
			Id:        song.Id,
			Title:     song.Title,
			Year:      song.Year,
			Performer: song.Performer,
			Genre:     song.Genre,
			Duration:  song.Duration,
		})
	}

	return web.AlbumResponse{
		Id:    album.Id,
		Name:  album.Name,
		Year:  album.Year,
		Songs: songs,
	}
}

func ToAlbumResponses(albums []domain.Album) []web.AlbumResponse {
	var albumResponses []web.AlbumResponse
	for _, album := range albums {
		albumResponses = append(albumResponses, ToAlbumResponse(album))
	}

	return albumResponses
}

func ToSongResponse(song domain.Song) web2.SongResponse {
	return web2.SongResponse{
		Id:        song.Id,
		Title:     song.Title,
		Year:      song.Year,
		Performer: song.Performer,
		Genre:     song.Genre,
		Duration:  song.Duration,
		AlbumId:   song.AlbumId,
	}
}

func ToSongResponses(songs []domain.Song) []web2.SongResponse {
	var SongResponses []web2.SongResponse
	for _, song := range songs {
		SongResponses = append(SongResponses, ToSongResponse(song))
	}

	return SongResponses
}

func ToUserRegisterResponse(user domain.User) web3.UserRegisterResponse {
	return web3.UserRegisterResponse{
		Id:       user.Id,
		Fullname: user.Fullname,
		Username: user.Username,
	}
}

func ToUserLoginResponse(token string) web3.UserLoginResponse {
	return web3.UserLoginResponse{
		Token: token,
	}
}
