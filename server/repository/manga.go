package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type Manga struct{ db *pgxpool.Pool }

func (r *Manga) GetOne(ctx context.Context, ID int32) (*model.Manga, error) {
	var m model.Manga

	err := r.db.QueryRow(
		ctx,
		"SELECT * FROM manga WHERE id = $1",
		ID,
	).Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo, &m.StartDate, &m.EndDate)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		log.Println("manga row scan failed:", err)
		return nil, err
	}

	return &m, err
}

func (r *Manga) GetAll(ctx context.Context) ([]*model.Manga, error) {
	return r.getList(r.db.Query(
		ctx,
		"SELECT * FROM manga ORDER BY name",
	))
}

func (r *Manga) GetByMangaka(ctx context.Context, mangaka *model.Mangaka) ([]*model.Manga, error) {
	return r.getList(r.db.Query(
		ctx,
		`
		SELECT m.*
		FROM manga m
		JOIN manga_mangaka_job mmkaj ON m.id = mmkaj.manga_id
		WHERE mmkaj.mangaka_id = $1
		ORDER BY name
		`,
		mangaka.ID,
	))
}

func (r *Manga) GetBySeriesMangaka(ctx context.Context, seriesMangaka *model.SeriesMangaka) ([]*model.Manga, error) {
	return r.getList(r.db.Query(
		ctx,
		`
		SELECT m.*
		FROM manga m
		JOIN manga_mangaka_job mmkaj ON m.id = mmkaj.manga_id
		WHERE mmkaj.mangaka_id = $1
		ORDER BY name
		`,
		seriesMangaka.ID,
	))
}

func (r *Manga) getList(rows pgx.Rows, err error) ([]*model.Manga, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*model.Manga

	for rows.Next() {
		var m model.Manga

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo, &m.StartDate, &m.EndDate)
		if err != nil {
			log.Println("Manga row scan failed:", err)
		}

		list = append(list, &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

func (r *Manga) GetGenres(ctx context.Context, manga *model.Manga) ([]string, error) {
	var list = []string{}

	rows, err := r.db.Query(
		ctx,
		`
			SELECT genre FROM manga_genre
			WHERE manga_id = $1
			ORDER BY genre
		`,
		manga.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var g string

		err := rows.Scan(&g)
		if err != nil {
			log.Println("Genre row scan failed:", err)
		}

		list = append(list, g)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

func NewManga(db *pgxpool.Pool) *Manga {
	return &Manga{db: db}
}
