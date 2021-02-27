package testhelpers

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"

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
