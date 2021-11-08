package repository

import (
	"context"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Chapter struct{ db *pgxpool.Pool }

func (r *Chapter) GetOne(ctx context.Context, mangaID int32, chapterNum float64) (*model.Chapter, error) {
	var c model.Chapter

	err := r.db.QueryRow(
		ctx,
		"SELECT * FROM chapter WHERE manga_id = $1 AND chapter_num = $2",
		mangaID, chapterNum,
	).Scan(&c.MangaID, &c.ChapterNum, &c.Name, &c.PageCount, &c.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		log.Println("Chapter row scan failed:", err)
		return nil, err
	}

	return &c, err
}

func NewChapter(db *pgxpool.Pool) *Chapter {
	return &Chapter{db: db}
}
