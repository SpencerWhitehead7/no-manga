package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type Loader struct {
	manga *dataloader.Loader
}

func (l *Loader) Manga(ctx context.Context, id int32) (*model.Manga, error) {
	v, err := l.manga.Load(ctx, int32Key(id))()
	if v == nil || err != nil {
		return nil, err
	}

	return v.(*model.Manga), nil
}

// func (l *Loader) checkResult TODO: generics :/

func NewLoader(db *pgxpool.Pool, shouldCache bool) *Loader {
	mangaBFs := newMangaBFs(db)

	var cache dataloader.Cache
	if !shouldCache {
		cache = &dataloader.NoCache{}
	}

	return &Loader{
		manga: dataloader.NewBatchedLoader(mangaBFs.byID, dataloader.WithCache(cache)),
	}
}
