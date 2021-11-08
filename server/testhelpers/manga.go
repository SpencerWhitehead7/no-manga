package testhelpers

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type MangaStub struct {
	Name        string
	OtherNames  []string
	Description string
	Demo        string
	StartDate   string
	EndDate     string
}

type MangaRow struct {
	ID          int
	Name        string
	OtherNames  *[]string
	Description string
	Demo        string
	StartDate   time.Time
	EndDate     *time.Time
}

func MangaFactory(t *testing.T, db *pgxpool.Pool, mangaStub MangaStub) MangaRow {
	var name string
	if mangaStub.Name != "" {
		name = mangaStub.Name
	} else {
		name = "tName"
	}
	var otherNames *[]string
	if len(mangaStub.OtherNames) != 0 {
		otherNames = &mangaStub.OtherNames
	} else {
		otherNames = nil
	}
	var description string
	if mangaStub.Description != "" {
		description = mangaStub.Description
	} else {
		description = "tDescription"
	}
	var demo string
	if mangaStub.Demo != "" {
		demo = mangaStub.Demo
	} else {
		demo = "shonen"
	}
	var startDate string
	if mangaStub.StartDate != "" {
		startDate = mangaStub.StartDate
	} else {
		startDate = "2000-01-01"
	}
	var endDate *string
	if mangaStub.EndDate != "" {
		endDate = &mangaStub.EndDate
	} else {
		endDate = nil
	}

	var m MangaRow

	err := db.QueryRow(
		context.Background(),
		`
		INSERT INTO manga(name, other_names, description, demo, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
		`,
		name, otherNames, description, demo, startDate, endDate,
	).Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo, &m.StartDate, &m.EndDate)
	if err != nil {
		t.Errorf("Failed to create manga row: %v", err)
	}

	return m
}
