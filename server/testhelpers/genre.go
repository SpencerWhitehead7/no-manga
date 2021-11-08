package testhelpers

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

type GenreStub struct {
	Name string
}

type GenreRow struct {
	Name string
}

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
