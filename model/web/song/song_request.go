package web

type CreateSong struct {
	Title     string `json:"title" validate:"required,min=1,max=255"`
	Year      int    `json:"year" validate:"required,min=1,max=255"`
	Genre     string `json:"genre" validate:"required,min=1,max=255"`
	Performer string `json:"performer" validate:"required,min=1,max=255"`
	Duration  int    `json:"duration" validate:"required,min=1,max=255"`
	AlbumId   int    `json:"albumId" validate:"min=1,max=255"`
}

type UpdateSong struct {
	Id        int    `json:"id"`
	Title     string `json:"title" validate:"required,min=1,max=255"`
	Year      int    `json:"year" validate:"required,min=1,max=255"`
	Genre     string `json:"genre" validate:"required,min=1,max=255"`
	Performer string `json:"performer" validate:"required,min=1,max=255"`
	Duration  int    `json:"duration" validate:"required,min=1,max=255"`
	AlbumId   int    `json:"albumId" validate:"min=1,max=255"`
}
