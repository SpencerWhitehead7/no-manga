package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type Loader struct {
	chapter             *dataloader.Loader
	chapterCount        *dataloader.Loader
	chapterList         *dataloader.Loader
	chapterListByManga  *dataloader.Loader
	genres              *dataloader.Loader
	magazine            *dataloader.Loader
	magazineList        *dataloader.Loader
	magazineListByManga *dataloader.Loader
	manga               *dataloader.Loader
	mangaList           *dataloader.Loader
	mangaListByMagazine *dataloader.Loader
	mangaListByMangaka  *dataloader.Loader
	mangaka             *dataloader.Loader
	mangakaList         *dataloader.Loader
	seriesMangakaList   *dataloader.Loader
}

func (l *Loader) Chapter(ctx context.Context, chapterID model.ChapterID) (*model.Chapter, error) {
	v, err := l.chapter.Load(ctx, chapterKey(chapterID))()
	return handleLoaderResult[*model.Chapter](v, err)
}

func (l *Loader) ChapterCount(ctx context.Context, manga *model.Manga) (int32, error) {
	v, err := l.chapterCount.Load(ctx, chapterCountKey(manga.ID))()
	return handleLoaderResult[int32](v, err)
}

func (l *Loader) ChapterList(ctx context.Context) ([]*model.Chapter, error) {
	v, err := l.chapterList.Load(ctx, dataloader.StringKey("chapterList"))()
	return handleLoaderResult[[]*model.Chapter](v, err)
}

func (l *Loader) ChapterListByManga(ctx context.Context, manga *model.Manga) ([]*model.Chapter, error) {
	v, err := l.chapterListByManga.Load(ctx, chapterListByMangaKey(manga.ID))()
	return handleLoaderResult[[]*model.Chapter](v, err)
}

func (l *Loader) Genres(ctx context.Context, manga *model.Manga) ([]string, error) {
	v, err := l.genres.Load(ctx, genresKey(manga.ID))()
	return handleLoaderResult[[]string](v, err)
}

func (l *Loader) Magazine(ctx context.Context, id int32) (*model.Magazine, error) {
	v, err := l.magazine.Load(ctx, magazineKey(id))()
	return handleLoaderResult[*model.Magazine](v, err)
}

func (l *Loader) MagazineList(ctx context.Context) ([]*model.Magazine, error) {
	v, err := l.magazineList.Load(ctx, dataloader.StringKey("magazineList"))()
	return handleLoaderResult[[]*model.Magazine](v, err)
}

func (l *Loader) MagazineListByManga(ctx context.Context, manga *model.Manga) ([]*model.Magazine, error) {
	v, err := l.magazineListByManga.Load(ctx, magazineListByMangaKey(manga.ID))()
	return handleLoaderResult[[]*model.Magazine](v, err)
}

func (l *Loader) Manga(ctx context.Context, id int32) (*model.Manga, error) {
	v, err := l.manga.Load(ctx, mangaKey(id))()
	return handleLoaderResult[*model.Manga](v, err)
}

func (l *Loader) MangaList(ctx context.Context) ([]*model.Manga, error) {
	v, err := l.mangaList.Load(ctx, dataloader.StringKey("mangaList"))()
	return handleLoaderResult[[]*model.Manga](v, err)
}

func (l *Loader) MangaListByMagazine(ctx context.Context, magazine *model.Magazine) ([]*model.Manga, error) {
	v, err := l.mangaListByMagazine.Load(ctx, mangaListByMagazineKey(magazine.ID))()
	return handleLoaderResult[[]*model.Manga](v, err)
}

func (l *Loader) MangaListByMangaka(ctx context.Context, mangaka *model.Mangaka) ([]*model.Manga, error) {
	v, err := l.mangaListByMangaka.Load(ctx, mangaListByMangakaKey(mangaka.ID))()
	return handleLoaderResult[[]*model.Manga](v, err)
}

func (l *Loader) MangaListBySeriesMangaka(ctx context.Context, seriesMangaka *model.SeriesMangaka) ([]*model.Manga, error) {
	v, err := l.mangaListByMangaka.Load(ctx, mangaListBySeriesMangakaKey(seriesMangaka.ID))()
	return handleLoaderResult[[]*model.Manga](v, err)
}

func (l *Loader) Mangaka(ctx context.Context, id int32) (*model.Mangaka, error) {
	v, err := l.mangaka.Load(ctx, mangakaKey(id))()
	return handleLoaderResult[*model.Mangaka](v, err)
}

func (l *Loader) MangakaList(ctx context.Context) ([]*model.Mangaka, error) {
	v, err := l.mangakaList.Load(ctx, dataloader.StringKey("mangakaList"))()
	return handleLoaderResult[[]*model.Mangaka](v, err)
}

func (l *Loader) SeriesMangakaList(ctx context.Context, manga *model.Manga) ([]*model.SeriesMangaka, error) {
	v, err := l.seriesMangakaList.Load(ctx, seriesMangakaListKey(manga.ID))()
	return handleLoaderResult[[]*model.SeriesMangaka](v, err)
}

func NewLoader(db *pgxpool.Pool, shouldCache bool) *Loader {
	chapterBFs := newChapterBFs(db)
	magazineBFs := newMagazineBFs(db)
	mangaBFs := newMangaBFs(db)
	mangakaBFs := newMangakaBFs(db)

	var cache dataloader.Cache
	if shouldCache {
		cache = newRistrettoCache()
	} else {
		cache = &dataloader.NoCache{}
	}

	return &Loader{
		chapter:             dataloader.NewBatchedLoader(chapterBFs.byID, dataloader.WithCache(cache)),
		chapterCount:        dataloader.NewBatchedLoader(chapterBFs.count, dataloader.WithCache(cache)),
		chapterList:         dataloader.NewBatchedLoader(chapterBFs.list, dataloader.WithCache(cache)),
		chapterListByManga:  dataloader.NewBatchedLoader(chapterBFs.listByManga, dataloader.WithCache(cache)),
		genres:              dataloader.NewBatchedLoader(mangaBFs.genres, dataloader.WithCache(cache)),
		magazine:            dataloader.NewBatchedLoader(magazineBFs.byID, dataloader.WithCache(cache)),
		magazineList:        dataloader.NewBatchedLoader(magazineBFs.list, dataloader.WithCache(cache)),
		magazineListByManga: dataloader.NewBatchedLoader(magazineBFs.listByManga, dataloader.WithCache(cache)),
		manga:               dataloader.NewBatchedLoader(mangaBFs.byID, dataloader.WithCache(cache)),
		mangaList:           dataloader.NewBatchedLoader(mangaBFs.list, dataloader.WithCache(cache)),
		mangaListByMagazine: dataloader.NewBatchedLoader(mangaBFs.listByMagazine, dataloader.WithCache(cache)),
		mangaListByMangaka:  dataloader.NewBatchedLoader(mangaBFs.listByMangaka, dataloader.WithCache(cache)),
		mangaka:             dataloader.NewBatchedLoader(mangakaBFs.byID, dataloader.WithCache(cache)),
		mangakaList:         dataloader.NewBatchedLoader(mangakaBFs.list, dataloader.WithCache(cache)),
		seriesMangakaList:   dataloader.NewBatchedLoader(mangakaBFs.listByManga, dataloader.WithCache(cache)),
	}
}

func handleLoaderResult[T any](result interface{}, err error) (T, error) {
	if result == nil || err != nil {
		var nilResult T
		return nilResult, err
	}
	return result.(T), nil
}
