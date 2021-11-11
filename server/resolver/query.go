package resolver

import (
	"context"

	"github.com/SpencerWhitehead7/no-manga/server/model"
	"github.com/SpencerWhitehead7/no-manga/server/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Query struct {
	chapterRepository *repository.Chapter
	mangakaRepository *repository.Mangaka
	mangaRepository   *repository.Manga
}

func (q *Query) Manga(
	ctx context.Context,
	args struct {
		ID int32
	},
) (*mangaResolver, error) {
	m, err := q.mangaRepository.GetOne(ctx, args.ID)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, nil
	}

	return &mangaResolver{manga: m, query: q}, nil
}

func (q *Query) MangaList(
	ctx context.Context,
) ([]*mangaResolver, error) {
	return q.mangaMListToRList(q.mangaRepository.GetAll(ctx))
}
func (q *Query) mangaMListToRList(mList []*model.Manga, err error) ([]*mangaResolver, error) {
	if err != nil {
		return nil, err
	}

	rList := make([]*mangaResolver, len(mList))
	for i, m := range mList {
		rList[i] = &mangaResolver{manga: m, query: q}
	}

	return rList, nil
}

func (q *Query) Chapter(
	ctx context.Context,
	args struct {
		MangaID    int32
		ChapterNum float64
	},
) (*chapterResolver, error) {
	c, err := q.chapterRepository.GetOne(ctx, args.MangaID, args.ChapterNum)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, nil
	}

	return &chapterResolver{chapter: c, query: q}, nil
}

func (q *Query) ChapterList(
	ctx context.Context,
) ([]*chapterResolver, error) {
	return q.chapterMListToRList(q.chapterRepository.GetAll(ctx))
}
func (q *Query) chapterMListToRList(mList []*model.Chapter, err error) ([]*chapterResolver, error) {
	if err != nil {
		return nil, err
	}

	rlist := make([]*chapterResolver, len(mList))
	for i, c := range mList {
		rlist[i] = &chapterResolver{chapter: c, query: q}
	}

	return rlist, nil
}

func (q *Query) Mangaka(
	ctx context.Context,
	args struct {
		ID int32
	},
) (*mangakaResolver, error) {
	m, err := q.mangakaRepository.GetOne(ctx, args.ID)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, nil
	}

	return &mangakaResolver{mangaka: m, query: q}, nil
}

func (q *Query) MangakaList(
	ctx context.Context,
) ([]*mangakaResolver, error) {
	mList, err := q.mangakaRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	rList := make([]*mangakaResolver, len(mList))
	for i, m := range mList {
		rList[i] = &mangakaResolver{mangaka: m, query: q}
	}

	return rList, nil
}

func (q *Query) seriesMangakaList(ctx context.Context, manga *model.Manga) ([]*seriesMangakaResolver, error) {
	sList, err := q.mangakaRepository.GetByManga(ctx, manga)
	if err != nil {
		return nil, err
	}

	rList := make([]*seriesMangakaResolver, len(sList))
	for i, s := range sList {
		rList[i] = &seriesMangakaResolver{seriesMangaka: s}
	}

	return rList, nil
}

func (q *Query) Magazine(
	ctx context.Context,
	args struct {
		ID int32
	},
) (*magazineResolver, error) {
	return &magazineResolver{}, nil
}

func (q *Query) MagazineList(
	ctx context.Context,
) ([]*magazineResolver, error) {
	return make([]*magazineResolver, 0), nil
}

func NewQuery(db *pgxpool.Pool) *Query {
	return &Query{
		chapterRepository: repository.NewChapter(db),
		mangakaRepository: repository.NewMangaka(db),
		mangaRepository:   repository.NewManga(db),
	}
}
