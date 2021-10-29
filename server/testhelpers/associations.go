package testhelpers

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

// MangaToGenres creates manga_genre rows between a manga row and genre rows.
func MangaToGenres(t *testing.T, db *pgxpool.Pool, manga MangaRow, genres []GenreRow) {
	query := "INSERT INTO manga_genre (manga_id, genre) VALUES"
	values := []interface{}{}
	i := 1
	for _, g := range genres {
		query += fmt.Sprintf(" ($%d, $%d),", i, i+1)
		values = append(values, manga.ID, g.Name)
		i += 2
	}
	if len(genres) != 0 {
		query = query[0 : len(query)-1]
	}

	_, err := db.Exec(context.Background(), query, values...)
	if err != nil {
		t.Errorf("Failed to create manga_genre row: %v", err)
	}
}

func MangaToMangaka(t *testing.T, db *pgxpool.Pool, manga MangaRow, mangaka MangakaRow, job string) {
	query := "INSERT INTO manga_mangaka_job (manga_id, mangaka_id, job) VALUES ($1, $2, $3)"
	values := []interface{}{manga.ID, mangaka.ID, job}

	_, err := db.Exec(context.Background(), query, values...)
	if err != nil {
		t.Errorf("Failed to create manga_mangaka_job row: %v", err)
	}
}

func MangaToMagazine(t *testing.T, db *pgxpool.Pool, manga MangaRow, magazine MagazineRow) {
	query := "INSERT INTO magazine_manga (manga_id, magazine_id) VALUES ($1, $2)"
	values := []interface{}{manga.ID, magazine.ID}

	_, err := db.Exec(context.Background(), query, values...)
	if err != nil {
		t.Errorf("Failed to create magazine_manga row: %v", err)
	}
}
