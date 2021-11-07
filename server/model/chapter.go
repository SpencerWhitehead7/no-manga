package model

import (
	"strconv"
	"time"
)

type Chapter struct {
	MangaID    int32     `json:"mangaId"`
	ChapterNum float64   `json:"chapterNum"`
	Name       *string   `json:"name"`
	PageCount  int32     `json:"pageCount"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Manga      *Manga    `json:"manga"`
}

func (c Chapter) ID() string {
	return strconv.FormatInt(int64(c.MangaID), 10) + "__" + strconv.FormatFloat(c.ChapterNum, 'f', -1, 64)
}
