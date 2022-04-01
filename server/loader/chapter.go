package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/model"
	"github.com/SpencerWhitehead7/no-manga/server/repository"
)

type chapterBFs struct{ chapterRepository *repository.Chapter }

// not using normal helpers because model.ChapterID keys are a special case
func (l *chapterBFs) byID(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := make([]model.ChapterID, len(keys))
	for i, k := range keys {
		ids[i] = k.Raw().(model.ChapterID)
	}

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

func (l *chapterBFs) count(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToCount, err := l.chapterRepository.GetCountByMangas(ctx, ids)

	return handleBatch(keys, ids, idToCount, err)
}

func (l *chapterBFs) list(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	cList, err := l.chapterRepository.GetAll(ctx)

	return handleSingle(keys, cList, err)
}

func (l *chapterBFs) listByManga(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := keysToIDs(keys)

	idToChapters, err := l.chapterRepository.GetByMangas(ctx, ids)

	return handleBatch(keys, ids, idToChapters, err)
}

func newChapterBFs(db *pgxpool.Pool) *chapterBFs {
	return &chapterBFs{chapterRepository: repository.NewChapter(db)}
}
