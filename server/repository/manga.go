package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type Manga struct{ db *pgxpool.Pool }

func (r *Manga) GetByIDs(ctx context.Context, ids []int32) (map[int32]*model.Manga, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT * FROM manga WHERE id = ANY($1)",
		ids,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idToManga := make(map[int32]*model.Manga, len(ids))

	for rows.Next() {
		var m model.Manga

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo, &m.StartDate, &m.EndDate)
		if err != nil {
			log.Println("Manga row scan failed:", err)
		}

		idToManga[m.ID] = &m
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return idToManga, nil
}

func (r *Manga) GetAll(ctx context.Context) ([]*model.Manga, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT * FROM manga ORDER BY name",
	)
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

func (r *Manga) GetByMangakas(ctx context.Context, ids []int32) (map[int32][]*model.Manga, error) {
	return r.getMap(r.db.Query(
		ctx,
		`
		SELECT m.*, mmkaj.mangaka_id
		FROM manga m
		JOIN manga_mangaka_job mmkaj ON m.id = mmkaj.manga_id
		WHERE mmkaj.mangaka_id = ANY($1)
		ORDER BY name
		`,
		ids,
	))
}

func (r *Manga) GetByMagazines(ctx context.Context, ids []int32) (map[int32][]*model.Manga, error) {
	return r.getMap(r.db.Query(
		ctx,
		`
		SELECT m.*, magm.magazine_id
		FROM manga m
		JOIN magazine_manga magm ON m.id = magm.manga_id
		WHERE magm.magazine_id = ANY($1)
		ORDER BY name 
		`,
		ids,
	))
}

func (r *Manga) getMap(rows pgx.Rows, err error) (map[int32][]*model.Manga, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idToMangas := make(map[int32][]*model.Manga)

	for rows.Next() {
		var id int32
		var m model.Manga

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo, &m.StartDate, &m.EndDate, &id)
		if err != nil {
			log.Println("Manga row scan failed:", err)
		}

		idToMangas[id] = append(idToMangas[id], &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return idToMangas, nil
}

func (r *Manga) GetGenresByMangas(ctx context.Context, ids []int32) (map[int32][]string, error) {
	rows, err := r.db.Query(
		ctx,
		`
		SELECT manga_id, ARRAY_AGG(genre ORDER BY genre) as genres
		FROM manga_genre
		WHERE manga_id = ANY($1)
		GROUP BY manga_id;
		`,
		ids,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idToGenres := make(map[int32][]string, len(ids))

	for rows.Next() {
		var id int32
		var genres []string

		err := rows.Scan(&id, &genres)
		if err != nil {
			log.Println("Genre row scan failed:", err)
		}

		idToGenres[id] = genres
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return idToGenres, nil
}

func NewManga(db *pgxpool.Pool) *Manga {
	return &Manga{db: db}
}
