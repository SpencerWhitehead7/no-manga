package testhelpers

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

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
