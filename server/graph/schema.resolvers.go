package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/SpencerWhitehead7/no-manga/server/graph/generated"
	"github.com/SpencerWhitehead7/no-manga/server/graph/model"
)

func (r *chapterResolver) Manga(ctx context.Context, chapter *model.Chapter) (*model.Manga, error) {
	return r.MangaRepo.GetOne(chapter.MangaID)
}

func (r *mangaResolver) Genres(ctx context.Context, manga *model.Manga) ([]string, error) {
	return r.MangaRepo.GetGenres(manga)
}

func (r *mangaResolver) ChapterCount(ctx context.Context, obj *model.Manga) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mangaResolver) ChapterList(ctx context.Context, obj *model.Manga) ([]*model.Chapter, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mangaResolver) MangakaList(ctx context.Context, obj *model.Manga) ([]*model.Mangaka, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mangaResolver) MagazineList(ctx context.Context, obj *model.Manga) ([]*model.Magazine, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Manga(ctx context.Context, ID int) (*model.Manga, error) {
	return r.MangaRepo.GetOne(ID)
}

func (r *queryResolver) MangaList(ctx context.Context) ([]*model.Manga, error) {
	return r.MangaRepo.GetAll()
}

func (r *queryResolver) Chapter(ctx context.Context, mangaID int, chapterNum float64) (*model.Chapter, error) {
	return r.ChapterRepo.GetOne(mangaID, chapterNum)
}

func (r *queryResolver) ChapterList(ctx context.Context) ([]*model.Chapter, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Mangaka(ctx context.Context, id int) (*model.Mangaka, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MangakaList(ctx context.Context) ([]*model.Mangaka, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Magazine(ctx context.Context, id int) (*model.Magazine, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MagazineList(ctx context.Context) ([]*model.Magazine, error) {
	panic(fmt.Errorf("not implemented"))
}

// Chapter returns generated.ChapterResolver implementation.
func (r *Resolver) Chapter() generated.ChapterResolver { return &chapterResolver{r} }

// Manga returns generated.MangaResolver implementation.
func (r *Resolver) Manga() generated.MangaResolver { return &mangaResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type chapterResolver struct{ *Resolver }
type mangaResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
