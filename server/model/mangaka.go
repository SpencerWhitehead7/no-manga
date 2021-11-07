package model

type Mangaka struct {
	ID          int32    `json:"id"`
	Name        string   `json:"name"`
	OtherNames  []string `json:"otherNames"`
	Description string   `json:"description"`
	MangaList   []*Manga `json:"mangaList"`
}
