package testhelpers

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

// GenreStub is the data necessary to create a manga.
type GenreStub struct {
	Name string
}

// GenreRow represents a row in the manga DB table.
type GenreRow struct {
	Name string
}

// GenreFactory creates a manga row from a GenreStub with defaults and returns it.
func GenreFactory(t *testing.T, db *pgxpool.Pool, genreStub GenreStub) GenreRow {
	var name string
	if genreStub.Name != "" {
		name = genreStub.Name
	} else {
		name = "tName"
	}

	var g GenreRow

	err := db.QueryRow(
		context.Background(),
		`
		INSERT INTO genre(name)
		VALUES ($1)
		RETURNING *
		`,
		name,
	).Scan(&g.Name)
	if err != nil {
		t.Errorf("Failed to create genre row: %v", err)
	}

	return g
}
