package model

import "time"

// Chapter represents a row in the chapter table.
type Chapter struct {
	ID         string    `json:"id"`
	MangaID    int       `json:"mangaId"`
	ChapterNum float64   `json:"chapterNum"`
	Name       *string   `json:"name"`
	PageCount  int       `json:"pageCount"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
