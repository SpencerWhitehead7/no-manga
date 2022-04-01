package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/repository"
)

type mangakaBFs struct{ mangakaRepository *repository.Mangaka }

func (l *mangakaBFs) byID(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToMangaka, err := l.mangakaRepository.GetByIDs(ctx, ids)

	return handleBatch(keys, ids, idToMangaka, err)
}

func (l *mangakaBFs) list(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	mList, err := l.mangakaRepository.GetAll(ctx)

	return handleSingle(keys, mList, err)
}

func (l *mangakaBFs) listByManga(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToSeriesMangakas, err := l.mangakaRepository.GetByMangas(ctx, ids)

	return handleBatch(keys, ids, idToSeriesMangakas, err)
}

func newMangakaBFs(db *pgxpool.Pool) *mangakaBFs {
	return &mangakaBFs{mangakaRepository: repository.NewMangaka(db)}
}
