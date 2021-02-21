package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// Manga is an interface for interacting with the DB to handle manga data.
type Manga struct{ db *pgxpool.Pool }

// MangaFactory creates new MangaRepositories.
func MangaFactory(db *pgxpool.Pool) *Manga {
	return &Manga{db: db}
}
