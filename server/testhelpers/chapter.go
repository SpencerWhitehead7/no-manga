package testhelpers

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ChapterStub struct {
	Manga      MangaRow
	ChapterNum float32
	Name       string
	PageCount  int
}

type ChapterRow struct {
	MangaID    int
	ChapterNum float32
	Name       *string
	PageCount  int
	UpdatedAt  time.Time
}

func ChapterFactory(t *testing.T, db *pgxpool.Pool, chapterStub ChapterStub) ChapterRow {
	var mangaID int
	if chapterStub.Manga.ID != 0 {
		mangaID = chapterStub.Manga.ID
	} else {
		mangaID = 1
	}
	var chapterNum float32
	if chapterStub.ChapterNum != 0 {
		chapterNum = chapterStub.ChapterNum
	} else {
		chapterNum = 1
	}
	var name *string
	if chapterStub.Name != "" {
		name = &chapterStub.Name
	} else {
		name = nil
	}
	var pageCount int
	if chapterStub.PageCount != 0 {
		pageCount = chapterStub.PageCount
	} else {
		pageCount = 1
	}

	var c ChapterRow

	err := db.QueryRow(
		context.Background(),
		`
		INSERT INTO chapter(manga_id, chapter_num, name, page_count)
		VALUES ($1, $2, $3, $4)
		RETURNING *
		`,
		mangaID, chapterNum, name, pageCount,
	).Scan(&c.MangaID, &c.ChapterNum, &c.Name, &c.PageCount, &c.UpdatedAt)
	if err != nil {
		t.Errorf("Failed to create manga row: %v", err)
	}

	return c
}
