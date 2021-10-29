package repository

import (
	"context"
	"fmt"
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

// GetAll returns all manga if given no parent, or all manga belonging to parent, sorted alphabetically by name.
func (r *Manga) GetAll(ctx context.Context, parent interface{}) ([]*model.Manga, error) {
	var list []*model.Manga

	var rows pgx.Rows
	var err error
	// todo: generics :/
	switch t := parent.(type) {
	case nil:
		rows, err = r.db.Query(
			ctx,
			"SELECT * FROM manga ORDER BY name",
		)
	case *model.Mangaka:
		// they're identical but go can't access t.ID properly if you combine the cases
		rows, err = r.db.Query(
			ctx,
			`
			SELECT m.*
			FROM manga m
			JOIN manga_mangaka_job mmkaj ON m.id = mmkaj.manga_id
			WHERE mmkaj.mangaka_id = $1
			ORDER BY name
			`,
			t.ID,
		)
	case *model.SeriesMangaka:
		// they're identical but go can't access t.ID properly if you combine the cases
		rows, err = r.db.Query(
			ctx,
			`
			SELECT m.*
			FROM manga m
			JOIN manga_mangaka_job mmkaj ON m.id = mmkaj.manga_id
			WHERE mmkaj.mangaka_id = $1
			ORDER BY name
			`,
			t.ID,
		)
	default:
		return nil, fmt.Errorf("could not resolve all manga because parent was unrecognized type")
	}
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

// GetGenres returns all a manga's genres, sorted alphabetically by name.
func (r *Manga) GetGenres(ctx context.Context, manga *model.Manga) ([]string, error) {
	var list = []string{}

	rows, err := r.db.Query(
		ctx,
		`
			SELECT genre FROM manga_genre
			WHERE manga_id = $1
			ORDER BY genre
		`,
		manga.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var g string

		err := rows.Scan(&g)
		if err != nil {
			log.Println("Genre row scan failed:", err)
		}

		list = append(list, g)
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
