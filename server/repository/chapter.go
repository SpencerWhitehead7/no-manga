package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Chapter struct{ db *pgxpool.Pool }

func (r *Chapter) GetByIDs(ctx context.Context, ids []model.ChapterID) (map[int32]map[float64]*model.Chapter, error) {
	query := "SELECT * FROM chapter"
	args := make([]interface{}, len(ids)*2)
	argNum := 0
	for i, id := range ids {
		if i != 0 {
			query += " OR "
		} else {
			query += " WHERE "
		}
		query += fmt.Sprintf("(manga_id = $%d AND chapter_num = $%d)", argNum+1, argNum+2)
		args[argNum] = id.MangaID
		args[argNum+1] = id.ChapterNum
		argNum += 2
	}

	rows, err := r.db.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idToChapter := make(map[int32]map[float64]*model.Chapter, len(ids))

	for rows.Next() {
		var c model.Chapter

		err = rows.Scan(&c.MangaID, &c.ChapterNum, &c.Name, &c.PageCount, &c.UpdatedAt)
		if err != nil {
			log.Println("Chapter row scan failed:", err)
		}

		_, ok := idToChapter[c.MangaID]
		if !ok {
			idToChapter[c.MangaID] = make(map[float64]*model.Chapter)
		}
		idToChapter[c.MangaID][c.ChapterNum] = &c

	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return idToChapter, err
}

func (r *Chapter) GetAll(ctx context.Context) ([]*model.Chapter, error) {
	return r.getList(r.db.Query(
		ctx,
		"SELECT * FROM chapter ORDER BY updated_at DESC, manga_id, chapter_num",
	))
}

func (r *Chapter) GetByManga(ctx context.Context, manga *model.Manga) ([]*model.Chapter, error) {
	return r.getList(r.db.Query(
		ctx,
		"SELECT * FROM chapter WHERE manga_id = $1 ORDER BY chapter_num",
		manga.ID,
	))
}

func (r *Chapter) getList(rows pgx.Rows, err error) ([]*model.Chapter, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*model.Chapter

	for rows.Next() {
		var c model.Chapter

		err := rows.Scan(&c.MangaID, &c.ChapterNum, &c.Name, &c.PageCount, &c.UpdatedAt)
		if err != nil {
			log.Println("Chapter row scan failed:", err)
		}

		list = append(list, &c)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

func (r *Chapter) GetCount(ctx context.Context, manga *model.Manga) (int32, error) {
	var count int32

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

func NewChapter(db *pgxpool.Pool) *Chapter {
	return &Chapter{db: db}
}
