package testhelpers

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

// MangakaStub is the data necessary to create a mangaka.
type MangakaStub struct {
	Name        string
	OtherNames  []string
	Description string
}

// MangakaRow represents a row in the mangaka DB table.
type MangakaRow struct {
	ID          int
	Name        string
	OtherNames  *[]string
	Description string
}

// MangakaFactory creates a mangaka row from a MangakaStub with defaults and returns it.
func MangakaFactory(t *testing.T, db *pgxpool.Pool, mangakaStub MangakaStub) MangakaRow {
	var name string
	if mangakaStub.Name != "" {
		name = mangakaStub.Name
	} else {
		name = "tName"
	}
	var otherNames *[]string
	if len(mangakaStub.OtherNames) != 0 {
		otherNames = &mangakaStub.OtherNames
	} else {
		otherNames = nil
	}
	var description string
	if mangakaStub.Description != "" {
		description = mangakaStub.Description
	} else {
		description = "tDescription"
	}

	var m MangakaRow

	err := db.QueryRow(
		context.Background(),
		`
		INSERT INTO mangaka(name, other_names, description)
		VALUES ($1, $2, $3)
		RETURNING *
		`,
		name, otherNames, description,
	).Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description)
	if err != nil {
		t.Errorf("Failed to create mangaka row: %v", err)
	}

	return m
}
