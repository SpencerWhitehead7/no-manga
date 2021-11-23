package repository

import (
	"context"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Mangaka struct{ db *pgxpool.Pool }

func (r *Mangaka) GetByIDs(ctx context.Context, ids []int32) (map[int32]*model.Mangaka, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT * FROM mangaka WHERE id = ANY($1)",
		ids,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idToMangaka := make(map[int32]*model.Mangaka, len(ids))

	for rows.Next() {
		var m model.Mangaka

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description)
		if err != nil {
			log.Println("Mangaka row scan failed:", err)
		}

		idToMangaka[m.ID] = &m
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return idToMangaka, err
}

func (r *Mangaka) GetAll(ctx context.Context) ([]*model.Mangaka, error) {
	var list []*model.Mangaka

	rows, err := r.db.Query(
		ctx,
		"SELECT * FROM mangaka ORDER BY name",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m model.Mangaka

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description)
		if err != nil {
			log.Println("Mangaka row scan failed:", err)
		}

		list = append(list, &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

func (r *Mangaka) GetByMangas(ctx context.Context, ids []int32) (map[int32][]*model.SeriesMangaka, error) {
	rows, err := r.db.Query(
		ctx,
		`
		SELECT mka.*, mmkaj.job, mmkaj.manga_id
		FROM mangaka mka
		JOIN manga_mangaka_job mmkaj ON mka.id = mmkaj.mangaka_id
		WHERE mmkaj.manga_id = ANY($1)
		ORDER BY name
		`,
		ids,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idToSeriesMangakas := make(map[int32][]*model.SeriesMangaka)

	for rows.Next() {
		var id int32
		var m model.SeriesMangaka

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Job, &id)
		if err != nil {
			log.Println("SeriesMangaka row scan failed:", err)
		}

		idToSeriesMangakas[id] = append(idToSeriesMangakas[id], &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return idToSeriesMangakas, nil
}

func NewMangaka(db *pgxpool.Pool) *Mangaka {
	return &Mangaka{db: db}
}
