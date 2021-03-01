package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/graph/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Chapter is an interface for interacting with the DB to handle chapter data.
type Chapter struct{ db *pgxpool.Pool }

// GetOne returns the chapter of the specified manga with the specified chapterNum.
func (r *Chapter) GetOne(ctx context.Context, mangaID int, chapterNum float64) (*model.Chapter, error) {
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

	c.ID = fmt.Sprintf("%v__%v", c.MangaID, c.ChapterNum)

	return &c, err
}

// ChapterFactory creates new ChapterRepositories.
func ChapterFactory(db *pgxpool.Pool) *Chapter {
	return &Chapter{db: db}
}
