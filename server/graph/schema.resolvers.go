package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/SpencerWhitehead7/no-manga/server/graph/generated"
	"github.com/SpencerWhitehead7/no-manga/server/graph/model"
)

func (r *queryResolver) Manga(ctx context.Context, id int) (*model.Manga, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MangaList(ctx context.Context) ([]*model.Manga, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Chapter(ctx context.Context, mangaID int, chapterNum float64) (*model.Chapter, error) {
	panic(fmt.Errorf("not implemented"))
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
