package repository

import (
	"context"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/graph/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

// SeriesMangaka is an interface for interacting with the DB to handle series mangaka data.
type SeriesMangaka struct{ db *pgxpool.Pool }

// GetAll returns all series mangaka associated with a manga with their job on that series
func (r *SeriesMangaka) GetAll(ctx context.Context, manga *model.Manga) ([]*model.SeriesMangaka, error) {
	var list []*model.SeriesMangaka

	rows, err := r.db.Query(
		ctx,
		`
		SELECT mka.*, mmkaj.job
		FROM mangaka mka
		JOIN manga_mangaka_job mmkaj ON mka.id = mmkaj.mangaka_id
		WHERE mmkaj.manga_id = $1
		ORDER BY name
		`,
		manga.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m model.SeriesMangaka

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Job)
		if err != nil {
			log.Println("SeriesMangaka row scan failed:", err)
		}

		list = append(list, &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

// SeriesMangakaFactory creates new SeriesMangakaRepositories.
func SeriesMangakaFactory(db *pgxpool.Pool) *SeriesMangaka {
	return &SeriesMangaka{db: db}
}
