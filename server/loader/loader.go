package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type Loader struct {
	chapter            *dataloader.Loader
	chapterCount       *dataloader.Loader
	chapterList        *dataloader.Loader
	chapterListByManga *dataloader.Loader
	genres             *dataloader.Loader
	manga              *dataloader.Loader
	mangaList          *dataloader.Loader
}

func (l *Loader) Chapter(ctx context.Context, chapterID model.ChapterID) (*model.Chapter, error) {
	v, err := l.chapter.Load(ctx, chapterKey(chapterID))()
	if v == nil || err != nil {
		return nil, err
	}

	return v.(*model.Chapter), nil
}

func (l *Loader) ChapterCount(ctx context.Context, manga *model.Manga) (int32, error) {
	v, err := l.chapterCount.Load(ctx, int32Key(manga.ID))()
	if v == nil || err != nil {
		return 0, err
	}

	return v.(int32), nil
}

func (l *Loader) ChapterList(ctx context.Context) ([]*model.Chapter, error) {
	v, err := l.chapterList.Load(ctx, dataloader.StringKey("chapterList"))()
	if v == nil || err != nil {
		return nil, err
	}

	return v.([]*model.Chapter), nil
}

func (l *Loader) ChapterListByManga(ctx context.Context, manga *model.Manga) ([]*model.Chapter, error) {
	v, err := l.chapterListByManga.Load(ctx, int32Key(manga.ID))()
	if v == nil || err != nil {
		return nil, err
	}

	return v.([]*model.Chapter), nil
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
	chapterBFs := newChapterBFs(db)
	mangaBFs := newMangaBFs(db)

	var cache dataloader.Cache
	if !shouldCache {
		cache = &dataloader.NoCache{}
	}

	return &Loader{
		chapter:            dataloader.NewBatchedLoader(chapterBFs.byID, dataloader.WithCache(cache)),
		chapterCount:       dataloader.NewBatchedLoader(chapterBFs.count, dataloader.WithCache(cache)),
		chapterList:        dataloader.NewBatchedLoader(chapterBFs.list, dataloader.WithCache(cache)),
		chapterListByManga: dataloader.NewBatchedLoader(chapterBFs.listByManga, dataloader.WithCache(cache)),
		genres:             dataloader.NewBatchedLoader(mangaBFs.genres, dataloader.WithCache(cache)),
		manga:              dataloader.NewBatchedLoader(mangaBFs.byID, dataloader.WithCache(cache)),
		mangaList:          dataloader.NewBatchedLoader(mangaBFs.list, dataloader.WithCache(cache)),
	}
}
