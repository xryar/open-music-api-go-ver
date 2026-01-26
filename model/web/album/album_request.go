package web

type CreateAlbum struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
	Year int    `json:"year" validate:"required,min=1"`
}

type UpdateAlbum struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,min=1,max=255"`
	Year int    `json:"year" validate:"required,min=1"`
}
