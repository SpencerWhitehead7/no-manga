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

// GetAll returns all chapters of a manga, sorted by chapterNum, or all chapters, sorted by updatedAt, if no manga is given.
func (r *Chapter) GetAll(ctx context.Context, manga *model.Manga) ([]*model.Chapter, error) {
	var list []*model.Chapter

	var rows pgx.Rows
	var err error
	if manga == nil {
		rows, err = r.db.Query(
			ctx,
			"SELECT * FROM chapter ORDER BY updated_at DESC, manga_id, chapter_num",
		)
	} else {
		rows, err = r.db.Query(
			ctx,
			"SELECT * FROM chapter WHERE manga_id = $1 ORDER BY chapter_num",
			manga.ID,
		)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c model.Chapter

		err := rows.Scan(&c.MangaID, &c.ChapterNum, &c.Name, &c.PageCount, &c.UpdatedAt)
		if err != nil {
			log.Println("Chapter row scan failed:", err)
		}

		c.ID = fmt.Sprintf("%v__%v", c.MangaID, c.ChapterNum)

		list = append(list, &c)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

// GetChapterCount returns number of chapters of the specified manga.
func (r *Chapter) GetChapterCount(ctx context.Context, manga *model.Manga) (int, error) {
	var count int

	err := r.db.QueryRow(
		ctx,
		"SELECT count(*) FROM chapter WHERE manga_id = $1",
		manga.ID,
	).Scan(&count)
	if err != nil {
		log.Println("ChapterCount row scan failed:", err)
		return count, err
	}

	return count, nil
}

// ChapterFactory creates new ChapterRepositories.
func ChapterFactory(db *pgxpool.Pool) *Chapter {
	return &Chapter{db: db}
}
