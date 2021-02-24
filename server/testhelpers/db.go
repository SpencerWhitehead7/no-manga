package testhelpers

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

// GetDbpool tries to connect to the integration test database.
func GetDbpool() *pgxpool.Pool {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("Unable to identify current directory (needed to load .env.test)")
	}
	basepath := filepath.Dir(file)
	err := godotenv.Load(filepath.Join(basepath, "../.env.test"))
	if err != nil {
		log.Fatalln("Unable to load test env variables:", err)
	}

	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

	return db
}

// ClearDB truncates all tables and resets ID counters.
func ClearDB(t *testing.T, db *pgxpool.Pool) {
	_, err := db.Exec(
		context.Background(),
		`
			TRUNCATE
			magazine,
			mangaka,
			manga,
			chapter,
			genre,
			magazine_manga,
			manga_mangaka_job,
			manga_genre
			RESTART IDENTITY
		`,
	)
	if err != nil {
		t.Errorf("Failed to clear DB: %v", err)
	}
}
