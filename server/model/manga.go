package model

import (
	"time"
)

type Manga struct {
	ID           int32            `json:"id"`
	Name         string           `json:"name"`
	OtherNames   []string         `json:"otherNames"`
	Description  string           `json:"description"`
	Demo         string           `json:"demo"`
	StartDate    time.Time        `json:"startDate"`
	EndDate      *time.Time       `json:"endDate"`
	Genres       []string         `json:"genres"`
	ChapterCount int32            `json:"chapterCount"`
	ChapterList  []*Chapter       `json:"chapterList"`
	MangakaList  []*SeriesMangaka `json:"mangakaList"`
	MagazineList []*Magazine      `json:"magazineList"`
}
