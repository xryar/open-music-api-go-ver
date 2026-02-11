package helper

import (
	"open-music-go/model/domain"
	web "open-music-go/model/web/album"
	web3 "open-music-go/model/web/playlist"
	web2 "open-music-go/model/web/song"
	web4 "open-music-go/model/web/user"
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

func ToPlaylistResponse(playlist domain.Playlist) web3.PlaylistResponse {
	songs := make([]web2.SongResponse, 0)

	for _, song := range playlist.Songs {
		songs = append(songs, web2.SongResponse{
			Id:        song.Id,
			Title:     song.Title,
			Year:      song.Year,
			Performer: song.Performer,
			Genre:     song.Genre,
			Duration:  song.Duration,
		})
	}

	return web3.PlaylistResponse{
		Id:    playlist.Id,
		Name:  playlist.Name,
		Owner: playlist.Owner,
		Songs: songs,
	}
}

func ToPlaylistResponses(playlists []domain.Playlist) []web3.PlaylistResponse {
	var playlistResponses []web3.PlaylistResponse
	for _, playlist := range playlists {
		playlistResponses = append(playlistResponses, ToPlaylistResponse(playlist))
	}

	return playlistResponses
}

func ToUserRegisterResponse(user domain.User) web4.UserRegisterResponse {
	return web4.UserRegisterResponse{
		Id:       user.Id,
		Fullname: user.Fullname,
		Username: user.Username,
	}
}

func ToUserLoginResponse(token string) web4.UserLoginResponse {
	return web4.UserLoginResponse{
		Token: token,
	}
}
