package repository

import (
	"context"
	"log"

	"github.com/SpencerWhitehead7/no-manga/server/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Magazine struct{ db *pgxpool.Pool }

func (r *Magazine) GetByIDs(ctx context.Context, ids []int32) (map[int32]*model.Magazine, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT * FROM magazine WHERE id = ANY($1)",
		ids,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idToMagazine := make(map[int32]*model.Magazine, len(ids))

	for rows.Next() {
		var m model.Magazine

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo)
		if err != nil {
			log.Println("Magazine row scan failed:", err)
		}
		idToMagazine[m.ID] = &m
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return idToMagazine, err
}

func (r *Magazine) GetAll(ctx context.Context) ([]*model.Magazine, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT * FROM magazine ORDER BY name",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*model.Magazine

	for rows.Next() {
		var m model.Magazine

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo)
		if err != nil {
			log.Println("Magazine row scan failed:", err)
		}

		list = append(list, &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return list, nil
}

func (r *Magazine) GetByMangas(ctx context.Context, ids []int32) (map[int32][]*model.Magazine, error) {
	rows, err := r.db.Query(
		ctx,
		`
		SELECT mag.*, magm.manga_id
		FROM magazine mag
		JOIN magazine_manga magm ON mag.id = magm.magazine_id
		WHERE magm.manga_id = ANY($1)
		ORDER BY name
		`,
		ids,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idToMagazines := make(map[int32][]*model.Magazine)

	for rows.Next() {
		var id int32
		var m model.Magazine

		err := rows.Scan(&m.ID, &m.Name, &m.OtherNames, &m.Description, &m.Demo, &id)
		if err != nil {
			log.Println("Magazine row scan failed:", err)
		}

		idToMagazines[id] = append(idToMagazines[id], &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return idToMagazines, nil
}

func NewMagazine(db *pgxpool.Pool) *Magazine {
	return &Magazine{db: db}
}
