package services

import (
	"context"
	web "open-music-go/model/web/playlist"
)

type PlaylistService interface {
	CreatePlaylist(ctx context.Context, request web.CreatePlaylistRequest) (web.PlaylistResponse, error)
	AddSongToPlaylist(ctx context.Context, request web.PlaylistSongRequest) error
	DeleteSongInPlaylist(ctx context.Context, request web.PlaylistSongRequest) error
	FindPlaylistByOwner(ctx context.Context) ([]web.PlaylistResponse, error)
	DeletePlaylist(ctx context.Context, id int) error
	FindPlaylistById(ctx context.Context, id int) (web.PlaylistResponse, error)
	FindAllPlaylists(ctx context.Context) ([]web.PlaylistResponse, error)
}
