package resolver

import (
	"context"

	"github.com/SpencerWhitehead7/no-manga/server/loader"
	"github.com/SpencerWhitehead7/no-manga/server/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Query struct {
	// ordinarily loaders are created and scoped per request because if
	// different users should get different things for the same request
	// caching / batching / permissioning will leak between requests
	// but this API is completely stateless and user independent,
	// so I'm using a singleton
	loader *loader.Loader
}

func (q *Query) Manga(
	ctx context.Context,
	args struct {
		ID int32
	},
) (*mangaResolver, error) {
	m, err := q.loader.Manga(ctx, args.ID)
	return handleSingle(q, newMangaResolver, m, err)
}

func (q *Query) MangaList(
	ctx context.Context,
) ([]*mangaResolver, error) {
	mList, err := q.loader.MangaList(ctx)
	return handleList(q, newMangaResolver, mList, err)
}

func (q *Query) Chapter(
	ctx context.Context,
	args struct {
		MangaID    int32
		ChapterNum float64
	},
) (*chapterResolver, error) {
	c, err := q.loader.Chapter(ctx, args)
	return handleSingle(q, newChapterResolver, c, err)
}

func (q *Query) ChapterList(
	ctx context.Context,
) ([]*chapterResolver, error) {
	cList, err := q.loader.ChapterList(ctx)
	return handleList(q, newChapterResolver, cList, err)
}

func (q *Query) Mangaka(
	ctx context.Context,
	args struct {
		ID int32
	},
) (*mangakaResolver, error) {
	m, err := q.loader.Mangaka(ctx, args.ID)
	return handleSingle(q, newMangakaResolver, m, err)
}

func (q *Query) MangakaList(
	ctx context.Context,
) ([]*mangakaResolver, error) {
	mList, err := q.loader.MangakaList(ctx)
	return handleList(q, newMangakaResolver, mList, err)
}

func (q *Query) seriesMangakaList(ctx context.Context, manga *model.Manga) ([]*seriesMangakaResolver, error) {
	mList, err := q.loader.SeriesMangakaList(ctx, manga)
	return handleList(q, newSeriesMangakaResolver, mList, err)
}

func (q *Query) Magazine(
	ctx context.Context,
	args struct {
		ID int32
	},
) (*magazineResolver, error) {
	m, err := q.loader.Magazine(ctx, args.ID)
	return handleSingle(q, newMagazineResolver, m, err)
}

func (q *Query) MagazineList(
	ctx context.Context,
) ([]*magazineResolver, error) {
	mList, err := q.loader.MagazineList(ctx)
	return handleList(q, newMagazineResolver, mList, err)
}

func NewQuery(db *pgxpool.Pool, shouldDataLoaderCache bool) *Query {
	return &Query{loader: loader.NewLoader(db, shouldDataLoaderCache)}
}

func handleSingle[M any, R any](q *Query, newResolver func(*M, *Query) *R, m *M, err error) (*R, error) {
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, nil
	}

	return newResolver(m, q), nil
}

func handleList[M any, R any](q *Query, newResolver func(*M, *Query) *R, mList []*M, err error) ([]*R, error) {
	if err != nil {
		return nil, err
	}

	rList := make([]*R, len(mList))
	for i, m := range mList {
		rList[i] = newResolver(m, q)
	}

	return rList, nil
}
