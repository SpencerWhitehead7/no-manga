package repository

import (
	"context"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/graph/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Manga is an interface for interacting with the DB to handle manga data.
type Manga struct{ db *pgxpool.Pool }

// GetOne returns the manga with the specified ID.
func (r *Manga) GetOne(ctx context.Context, ID int) (*model.Manga, error) {
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

		log.Println("Manga row scan failed:", err)
		return nil, err
	}

	return &m, err
}

// MangaFactory creates new MangaRepositories.
func MangaFactory(db *pgxpool.Pool) *Manga {
	return &Manga{db: db}
}
