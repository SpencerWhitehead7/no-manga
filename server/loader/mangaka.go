package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/repository"
)

type mangakaBFs struct{ mangakaRepository *repository.Mangaka }

func (l *mangakaBFs) byID(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := int32KeysToIDs(keys)

	idToMangaka, err := l.mangakaRepository.GetByIDs(ctx, ids)
	if err != nil {
		return loadBatchError(keys, err)
	}

	loadBatchSuccess := make([]*dataloader.Result, len(ids))
	for i, id := range ids {
		loadBatchSuccess[i] = &dataloader.Result{Data: idToMangaka[id]}
	}

	return loadBatchSuccess
}

func newMangakaBFs(db *pgxpool.Pool) *mangakaBFs {
	return &mangakaBFs{mangakaRepository: repository.NewMangaka(db)}
}
