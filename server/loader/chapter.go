package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/repository"
)

type chapterBFs struct{ chapterRepository *repository.Chapter }

func (l *chapterBFs) byID(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := chapterKeysToIDs(keys)

	idToChapter, err := l.chapterRepository.GetByIDs(ctx, ids)
	if err != nil {
		return loadBatchError(keys, err)
	}

	loadBatchSuccess := make([]*dataloader.Result, len(keys))
	for i, id := range ids {
		loadBatchSuccess[i] = &dataloader.Result{Data: idToChapter[id.MangaID][id.ChapterNum]}
	}

	return loadBatchSuccess
}

func newChapterBFs(db *pgxpool.Pool) *chapterBFs {
	return &chapterBFs{chapterRepository: repository.NewChapter(db)}
}
