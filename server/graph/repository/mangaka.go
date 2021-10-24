package repository

import (
	"context"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/graph/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Mangaka is an interface for interacting with the DB to handle mangaka data.
type Mangaka struct{ db *pgxpool.Pool }

// GetOne returns the mangaka with the specified ID.
func (r *Mangaka) GetOne(ctx context.Context, ID int) (*model.Mangaka, error) {
	var m model.Mangaka

	err := r.db.QueryRow(
		ctx,
		"SELECT * FROM mangaka WHERE id = $1",
		ID,
	).Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		log.Println("Mangaka row scan failed:", err)
		return nil, err
	}

	return &m, err
}

// GetAll returns all mangaka, sorted alphabetically by name.
func (r *Mangaka) GetAll(ctx context.Context) ([]*model.Mangaka, error) {
	var list []*model.Mangaka

	rows, err := r.db.Query(
		ctx,
		"SELECT * FROM mangaka ORDER BY name",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m model.Mangaka

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description)
		if err != nil {
			log.Println("Mangaka row scan failed:", err)
		}

		list = append(list, &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

// MangaFactory creates new MangaRepositories.
func MangakaFactory(db *pgxpool.Pool) *Mangaka {
	return &Mangaka{db: db}
}
