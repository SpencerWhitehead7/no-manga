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
	return getID(c.MangaID, c.ChapterNum)
}

type ChapterID struct {
	MangaID    int32
	ChapterNum float64
}

func (c ChapterID) ID() string {
	return getID(c.MangaID, c.ChapterNum)
}

func getID(mangaID int32, chapterNum float64) string {
	return strconv.FormatInt(int64(mangaID), 10) + "__" + strconv.FormatFloat(chapterNum, 'f', -1, 64)
}
