package testhelpers

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

type MangakaStub struct {
	Name        string
	OtherNames  []string
	Description string
}

type MangakaRow struct {
	ID          int
	Name        string
	OtherNames  *[]string
	Description string
}

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
