package resolver

import "context"

type Query struct{}

func (q *Query) Manga(
	ctx context.Context,
	args struct {
		ID int32
	},
) (*mangaResolver, error) {
	return &mangaResolver{}, nil
}

func (q *Query) MangaList(
	ctx context.Context,
) ([]*mangaResolver, error) {
	return make([]*mangaResolver, 0), nil
}

func (q *Query) Chapter(
	ctx context.Context,
	args struct {
		MangaID    int32
		ChapterNum float64
	},
) (*chapterResolver, error) {
	return &chapterResolver{}, nil
}

func (q *Query) ChapterList(
	ctx context.Context,
) ([]*chapterResolver, error) {
	return make([]*chapterResolver, 0), nil
}

func (q *Query) Mangaka(
	ctx context.Context,
	args struct {
		ID int32
	},
) (*mangakaResolver, error) {
	return &mangakaResolver{}, nil
}

func (q *Query) MangakaList(
	ctx context.Context,
) ([]*mangakaResolver, error) {
	return make([]*mangakaResolver, 0), nil
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

func NewQuery() *Query {
	return &Query{}
}
