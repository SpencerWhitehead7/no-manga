package testhelpers

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

// MagazineStub is the data necessary to create a magazine.
type MagazineStub struct {
	Name        string
	OtherNames  []string
	Description string
	Demo        string
}

// MagazineRow represents a row in the magazine DB table.
type MagazineRow struct {
	ID          int
	Name        string
	OtherNames  *[]string
	Description string
	Demo        string
}

// MagazineFactory creates a magazine row from a MagazineStub with defaults and returns it.
func MagazineFactory(t *testing.T, db *pgxpool.Pool, magazineStub MagazineStub) MagazineRow {
	var name string
	if magazineStub.Name != "" {
		name = magazineStub.Name
	} else {
		name = "tName"
	}
	var otherNames *[]string
	if len(magazineStub.OtherNames) != 0 {
		otherNames = &magazineStub.OtherNames
	} else {
		otherNames = nil
	}
	var description string
	if magazineStub.Description != "" {
		description = magazineStub.Description
	} else {
		description = "tDescription"
	}
	var demo string
	if magazineStub.Demo != "" {
		demo = magazineStub.Description
	} else {
		demo = "shonen"
	}

	var m MagazineRow

	err := db.QueryRow(
		context.Background(),
		`
		INSERT INTO magazine(name, other_names, description, demo)
		VALUES ($1, $2, $3, $4)
		RETURNING *
		`,
		name, otherNames, description, demo,
	).Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo)
	if err != nil {
		t.Errorf("Failed to create magazine row: %v", err)
	}

	return m
}
