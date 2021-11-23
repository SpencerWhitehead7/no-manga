package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/repository"
)

type mangaBFs struct{ mangaRepository *repository.Manga }

func (l *mangaBFs) byID(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := int32KeysToIDs(keys)

	idToManga, err := l.mangaRepository.GetByIDs(ctx, ids)
	if err != nil {
		return loadBatchError(keys, err)
	}

	loadBatchSuccess := make([]*dataloader.Result, len(ids))
	for i, id := range ids {
		loadBatchSuccess[i] = &dataloader.Result{Data: idToManga[id]}
	}

	return loadBatchSuccess
}

func (l *mangaBFs) list(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	mList, err := l.mangaRepository.GetAll(ctx)

	return handleSingleBatch(keys, mList, err)
}

func (l *mangaBFs) listByMangaka(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := int32KeysToIDs(keys)

	idToManga, err := l.mangaRepository.GetByMangakas(ctx, ids)
	if err != nil {
		return loadBatchError(keys, err)
	}

	loadBatchSuccess := make([]*dataloader.Result, len(ids))
	for i, id := range ids {
		loadBatchSuccess[i] = &dataloader.Result{Data: idToManga[id]}
	}

	return loadBatchSuccess
}

func (l *mangaBFs) genres(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := int32KeysToIDs(keys)

	idToGenres, err := l.mangaRepository.GetGenresByMangas(ctx, ids)
	if err != nil {
		return loadBatchError(keys, err)
	}

	loadBatchSuccess := make([]*dataloader.Result, len(ids))
	for i, id := range ids {
		loadBatchSuccess[i] = &dataloader.Result{Data: idToGenres[id]}
	}

	return loadBatchSuccess
}

func newMangaBFs(db *pgxpool.Pool) *mangaBFs {
	return &mangaBFs{mangaRepository: repository.NewManga(db)}
}
