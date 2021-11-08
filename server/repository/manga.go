package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type Manga struct{ db *pgxpool.Pool }

func (r *Manga) GetOne(ctx context.Context, ID int32) (*model.Manga, error) {
	var m model.Manga

	err := r.db.QueryRow(
		ctx,
		"SELECT * FROM manga WHERE id = $1",
		ID,
	).Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo, &m.StartDate, &m.EndDate)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		log.Println("manga row scan failed:", err)
		return nil, err
	}

	return &m, err
}

func NewManga(db *pgxpool.Pool) *Manga {
	return &Manga{db: db}
}
