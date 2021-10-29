package repository

import (
	"context"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/graph/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Magazine is an interface for interacting with the DB to handle magazine data.
type Magazine struct{ db *pgxpool.Pool }

// GetOne returns the magazine with the specified ID.
func (r *Magazine) GetOne(ctx context.Context, ID int) (*model.Magazine, error) {
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

// MagazineFactory creates new MagazineRepositories.
func MagazineFactory(db *pgxpool.Pool) *Magazine {
	return &Magazine{db: db}
}
