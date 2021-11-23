package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/repository"
)

type magazineBFs struct{ magazineRepository *repository.Magazine }

func (l *magazineBFs) byID(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := int32KeysToIDs(keys)

	idToMagazine, err := l.magazineRepository.GetByIDs(ctx, ids)
	if err != nil {
		return loadBatchError(keys, err)
	}

	loadBatchSuccess := make([]*dataloader.Result, len(ids))
	for i, id := range ids {
		loadBatchSuccess[i] = &dataloader.Result{Data: idToMagazine[id]}
	}

	return loadBatchSuccess
}

func (l *magazineBFs) list(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	mList, err := l.magazineRepository.GetAll(ctx)

	return handleSingleBatch(keys, mList, err)
}

func newMagazineBFs(db *pgxpool.Pool) *magazineBFs {
	return &magazineBFs{magazineRepository: repository.NewMagazine(db)}
}
