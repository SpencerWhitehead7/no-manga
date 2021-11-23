package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type Loader struct {
	genres    *dataloader.Loader
	manga     *dataloader.Loader
	mangaList *dataloader.Loader
}

func (l *Loader) Genres(ctx context.Context, manga *model.Manga) ([]string, error) {
	v, err := l.genres.Load(ctx, int32Key(manga.ID))()
	if v == nil || err != nil {
		return nil, err
	}

	return v.([]string), nil
}

func (l *Loader) Manga(ctx context.Context, id int32) (*model.Manga, error) {
	v, err := l.manga.Load(ctx, int32Key(id))()
	if v == nil || err != nil {
		return nil, err
	}

	return v.(*model.Manga), nil
}

func (l *Loader) MangaList(ctx context.Context) ([]*model.Manga, error) {
	v, err := l.mangaList.Load(ctx, dataloader.StringKey("mangaList"))()
	if v == nil || err != nil {
		return nil, err
	}

	return v.([]*model.Manga), nil
}

// func (l *Loader) checkResult TODO: generics :/

func NewLoader(db *pgxpool.Pool, shouldCache bool) *Loader {
	mangaBFs := newMangaBFs(db)

	var cache dataloader.Cache
	if !shouldCache {
		cache = &dataloader.NoCache{}
	}

	return &Loader{
		genres:    dataloader.NewBatchedLoader(mangaBFs.genres, dataloader.WithCache(cache)),
		manga:     dataloader.NewBatchedLoader(mangaBFs.byID, dataloader.WithCache(cache)),
		mangaList: dataloader.NewBatchedLoader(mangaBFs.list, dataloader.WithCache(cache)),
	}
}
