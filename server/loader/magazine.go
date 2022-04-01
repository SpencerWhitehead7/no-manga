package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/repository"
)

type magazineBFs struct{ magazineRepository *repository.Magazine }

func (l *magazineBFs) byID(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToMagazine, err := l.magazineRepository.GetByIDs(ctx, ids)

	return handleBatch(keys, ids, idToMagazine, err)
}

func (l *magazineBFs) list(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	mList, err := l.magazineRepository.GetAll(ctx)

	return handleSingle(keys, mList, err)
}

func (l *magazineBFs) listByManga(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToMagazines, err := l.magazineRepository.GetByMangas(ctx, ids)

	return handleBatch(keys, ids, idToMagazines, err)
}

func newMagazineBFs(db *pgxpool.Pool) *magazineBFs {
	return &magazineBFs{magazineRepository: repository.NewMagazine(db)}
}
