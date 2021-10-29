package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/SpencerWhitehead7/no-manga/server/graph/generated"
	"github.com/SpencerWhitehead7/no-manga/server/graph/model"
)

func (r *chapterResolver) Manga(ctx context.Context, obj *model.Chapter) (*model.Manga, error) {
	return r.MangaRepo.GetOne(ctx, obj.MangaID)
}

func (r *magazineResolver) MangaList(ctx context.Context, obj *model.Magazine) ([]*model.Manga, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mangaResolver) Genres(ctx context.Context, obj *model.Manga) ([]string, error) {
	return r.MangaRepo.GetGenres(ctx, obj)
}

func (r *mangaResolver) ChapterCount(ctx context.Context, obj *model.Manga) (int, error) {
	return r.ChapterRepo.GetChapterCount(ctx, obj)
}

func (r *mangaResolver) ChapterList(ctx context.Context, obj *model.Manga) ([]*model.Chapter, error) {
	return r.ChapterRepo.GetAll(ctx, obj)
}

func (r *mangaResolver) MangakaList(ctx context.Context, obj *model.Manga) ([]*model.SeriesMangaka, error) {
	return r.SeriesMangakaRepo.GetAll(ctx, obj)
}

func (r *mangaResolver) MagazineList(ctx context.Context, obj *model.Manga) ([]*model.Magazine, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mangakaResolver) MangaList(ctx context.Context, obj *model.Mangaka) ([]*model.Manga, error) {
	return r.MangaRepo.GetAll(ctx, obj)
}

func (r *queryResolver) Manga(ctx context.Context, id int) (*model.Manga, error) {
	return r.MangaRepo.GetOne(ctx, id)
}

func (r *queryResolver) MangaList(ctx context.Context) ([]*model.Manga, error) {
	return r.MangaRepo.GetAll(ctx, nil)
}

func (r *queryResolver) Chapter(ctx context.Context, mangaID int, chapterNum float64) (*model.Chapter, error) {
	return r.ChapterRepo.GetOne(ctx, mangaID, chapterNum)
}

func (r *queryResolver) ChapterList(ctx context.Context) ([]*model.Chapter, error) {
	return r.ChapterRepo.GetAll(ctx, nil)
}

func (r *queryResolver) Mangaka(ctx context.Context, id int) (*model.Mangaka, error) {
	return r.MangakaRepo.GetOne(ctx, id)
}

func (r *queryResolver) MangakaList(ctx context.Context) ([]*model.Mangaka, error) {
	return r.MangakaRepo.GetAll(ctx)
}

func (r *queryResolver) Magazine(ctx context.Context, id int) (*model.Magazine, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MagazineList(ctx context.Context) ([]*model.Magazine, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *seriesMangakaResolver) MangaList(ctx context.Context, obj *model.SeriesMangaka) ([]*model.Manga, error) {
	return r.MangaRepo.GetAll(ctx, obj)
}

// Chapter returns generated.ChapterResolver implementation.
func (r *Resolver) Chapter() generated.ChapterResolver { return &chapterResolver{r} }

// Magazine returns generated.MagazineResolver implementation.
func (r *Resolver) Magazine() generated.MagazineResolver { return &magazineResolver{r} }

// Manga returns generated.MangaResolver implementation.
func (r *Resolver) Manga() generated.MangaResolver { return &mangaResolver{r} }

// Mangaka returns generated.MangakaResolver implementation.
func (r *Resolver) Mangaka() generated.MangakaResolver { return &mangakaResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SeriesMangaka returns generated.SeriesMangakaResolver implementation.
func (r *Resolver) SeriesMangaka() generated.SeriesMangakaResolver { return &seriesMangakaResolver{r} }

type chapterResolver struct{ *Resolver }
type magazineResolver struct{ *Resolver }
type mangaResolver struct{ *Resolver }
type mangakaResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type seriesMangakaResolver struct{ *Resolver }
