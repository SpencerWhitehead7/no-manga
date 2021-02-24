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

// GetAll returns all manga, sorted alphabetically by name.
func (r *Manga) GetAll(ctx context.Context) ([]*model.Manga, error) {
	var list []*model.Manga

	rows, err := r.db.Query(
		ctx,
		"SELECT * FROM manga ORDER BY name",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m model.Manga

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo, &m.StartDate, &m.EndDate)
		if err != nil {
			log.Println("Manga row scan failed:", err)
		}

		list = append(list, &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

// MangaFactory creates new MangaRepositories.
func MangaFactory(db *pgxpool.Pool) *Manga {
	return &Manga{db: db}
}
