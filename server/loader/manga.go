package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/repository"
)

type mangaBFs struct{ mangaRepository *repository.Manga }

func (l *mangaBFs) byID(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToManga, err := l.mangaRepository.GetByIDs(ctx, ids)

	return handleBatch(keys, ids, idToManga, err)
}

func (l *mangaBFs) list(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	mList, err := l.mangaRepository.GetAll(ctx)

	return handleSingle(keys, mList, err)
}

func (l *mangaBFs) listByMagazine(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToManga, err := l.mangaRepository.GetByMagazines(ctx, ids)

	return handleBatch(keys, ids, idToManga, err)
}

func (l *mangaBFs) listByMangaka(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToManga, err := l.mangaRepository.GetByMangakas(ctx, ids)

	return handleBatch(keys, ids, idToManga, err)
}

func (l *mangaBFs) genres(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToGenres, err := l.mangaRepository.GetGenresByMangas(ctx, ids)

	return handleBatch(keys, ids, idToGenres, err)
}

func newMangaBFs(db *pgxpool.Pool) *mangaBFs {
	return &mangaBFs{mangaRepository: repository.NewManga(db)}
}
