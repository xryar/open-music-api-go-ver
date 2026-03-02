//go:build wireinject
// +build wireinject

package app

import (
	"database/sql"

	albumController "open-music-go/controllers/album"
	playlistController "open-music-go/controllers/playlist"
	collabController "open-music-go/controllers/playlist_collab"
	songController "open-music-go/controllers/song"
	userController "open-music-go/controllers/user"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"

	albumService "open-music-go/services/album"
	playlistService "open-music-go/services/playlist"
	collabService "open-music-go/services/playlist_collab"
	songService "open-music-go/services/song"
	userService "open-music-go/services/user"

	albumRepo "open-music-go/repositories/album"
	playlistRepo "open-music-go/repositories/playlist"
	activityRepo "open-music-go/repositories/playlist_activity"
	collabRepo "open-music-go/repositories/playlist_collab"
	songRepo "open-music-go/repositories/song"
	userRepo "open-music-go/repositories/user"
)

func ProvideDB() *sql.DB {
	return NewDB()
}

func ProvideValidator() *validator.Validate {
	return validator.New()
}

var repoSet = wire.NewSet(
	albumRepo.NewAlbumRepository,
	wire.Bind(new(albumRepo.AlbumRepository), new(*albumRepo.AlbumRepositoryImpl)),
	songRepo.NewSongRepository,
	wire.Bind(new(songRepo.SongRepository), new(*songRepo.SongRepositoryImpl)),
	userRepo.NewUserRepository,
	wire.Bind(new(userRepo.UserRepository), new(*userRepo.UserRepositoryImpl)),
	playlistRepo.NewPlaylistRepository,
	wire.Bind(new(playlistRepo.PlaylistRepository), new(*playlistRepo.PlaylistRepositoryImpl)),
	collabRepo.NewPlaylistCollabRepository,
	wire.Bind(new(collabRepo.PlaylistCollabRepository), new(*collabRepo.PlaylistCollabRepositoryImpl)),
	activityRepo.NewPlaylistActivityRepository,
	wire.Bind(new(activityRepo.PlaylistActivityRepository), new(*activityRepo.PlaylistActivityRepositoryImpl)),
)

var serviceSet = wire.NewSet(
	albumService.NewAlbumService,
	wire.Bind(new(albumService.AlbumService), new(*albumService.AlbumServiceImpl)),
	songService.NewSongService,
	wire.Bind(new(songService.SongService), new(*songService.SongServiceImpl)),
	userService.NewUserService,
	wire.Bind(new(userService.UserService), new(*userService.UserServiceImpl)),

	playlistService.NewPlaylistService,
	wire.Bind(new(playlistService.PlaylistService), new(*playlistService.PlaylistServiceImpl)),
	collabService.NewPlaylistCollabService,
	wire.Bind(new(collabService.PlaylistCollabService), new(*collabService.PlaylistCollabServiceImpl)),
)

var controllerSet = wire.NewSet(
	albumController.NewAlbumController,
	wire.Bind(new(albumController.AlbumController), new(*albumController.AlbumControllerImpl)),
	songController.NewSongController,
	wire.Bind(new(songController.SongController), new(*songController.SongControllerImpl)),
	userController.NewUserController,
	wire.Bind(new(userController.UserController), new(*userController.UserControllerImpl)),
	playlistController.NewPlaylistController,
	wire.Bind(new(playlistController.PlaylistController), new(*playlistController.PlaylistControllerImpl)),
	collabController.NewPlaylistCollabController,
	wire.Bind(new(collabController.PlaylistCollabController), new(*collabController.PlaylistCollabControllerImpl)),
)

func InitializeRouter() (*httprouter.Router, error) {
	wire.Build(
		ProvideDB,
		ProvideValidator,
		repoSet,
		serviceSet,
		controllerSet,
		NewRouter,
	)

	return nil, nil
}
