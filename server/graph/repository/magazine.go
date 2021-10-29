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

// GetAll returns all magazines a manga ran in, or all magazines if no manga is given, sorted alphabetically by name.
func (r *Magazine) GetAll(ctx context.Context, manga *model.Manga) ([]*model.Magazine, error) {
	var list []*model.Magazine

	var rows pgx.Rows
	var err error
	if manga == nil {
		rows, err = r.db.Query(
			ctx,
			"SELECT * FROM magazine ORDER BY name",
		)
	} else {
		rows, err = r.db.Query(
			ctx,
			`
			SELECT mag.*
			FROM magazine mag
			JOIN magazine_manga magm ON mag.id = magm.magazine_id
			WHERE magm.manga_id = $1
			ORDER BY name
			`,
			manga.ID,
		)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

// MagazineFactory creates new MagazineRepositories.
func MagazineFactory(db *pgxpool.Pool) *Magazine {
	return &Magazine{db: db}
}
