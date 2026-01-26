package helper

import (
	"open-music-go/model/domain"
	web "open-music-go/model/web/album"
)

func ToAlbumResponse(album domain.Album) web.AlbumResponse {
	return web.AlbumResponse{
		Id:   album.Id,
		Name: album.Name,
		Year: album.Year,
	}
}

func ToAlbumResponses(albums []domain.Album) []web.AlbumResponse {
	var albumResponses []web.AlbumResponse
	for _, album := range albums {
		albumResponses = append(albumResponses, ToAlbumResponse(album))
	}

	return albumResponses
}
