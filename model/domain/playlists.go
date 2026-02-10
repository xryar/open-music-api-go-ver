package domain

type Playlist struct {
	Id    int
	Name  string
	Owner int
	Songs []Song
}
