package web

type CreateAlbumRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
	Year int    `json:"year" validate:"required,min=1"`
}

type UpdateAlbumRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,min=1,max=255"`
	Year int    `json:"year" validate:"required,min=1"`
}
