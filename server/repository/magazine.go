package repository

import (
	"context"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Magazine struct{ db *pgxpool.Pool }

func (r *Magazine) GetOne(ctx context.Context, ID int32) (*model.Magazine, error) {
	var m model.Magazine

	err := r.db.QueryRow(
		ctx,
		"SELECT * FROM magazine WHERE id = $1",
		ID,
	).Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		log.Println("Magazine row scan failed:", err)
		return nil, err
	}

	return &m, err
}

func (r *Magazine) GetAll(ctx context.Context) ([]*model.Magazine, error) {
	return r.getList(r.db.Query(
		ctx,
		"SELECT * FROM magazine ORDER BY name",
	))
}

func (r *Magazine) GetByManga(ctx context.Context, manga *model.Manga) ([]*model.Magazine, error) {
	return r.getList(r.db.Query(
		ctx,
		`
		SELECT mag.*
		FROM magazine mag
		JOIN magazine_manga magm ON mag.id = magm.magazine_id
		WHERE magm.manga_id = $1
		ORDER BY name
		`,
		manga.ID,
	))
}

func (r *Magazine) getList(rows pgx.Rows, err error) ([]*model.Magazine, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*model.Magazine

	for rows.Next() {
		var m model.Magazine

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo)
		if err != nil {
			log.Println("Magazine row scan failed:", err)
		}

		list = append(list, &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

func NewMagazine(db *pgxpool.Pool) *Magazine {
	return &Magazine{db: db}
}
