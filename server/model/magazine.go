package model

type Magazine struct {
	ID          int32    `json:"id"`
	Name        string   `json:"name"`
	OtherNames  []string `json:"otherNames"`
	Description string   `json:"description"`
	Demo        string   `json:"demo"`
	MangaList   []*Manga `json:"mangaList"`
}
