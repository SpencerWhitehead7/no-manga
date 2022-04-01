package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type chapterResolver struct {
	chapter *model.Chapter
	query   *Query
}

func (r *chapterResolver) ID() graphql.ID {
	return graphql.ID(r.chapter.ID())
}
func (r *chapterResolver) ChapterNum() float64 {
	return r.chapter.ChapterNum
}
func (r *chapterResolver) Name() *string {
	return r.chapter.Name
}
func (r *chapterResolver) PageCount() int32 {
	return r.chapter.PageCount
}
func (r *chapterResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.chapter.UpdatedAt}
}
func (r *chapterResolver) Manga(ctx context.Context) (*mangaResolver, error) {
	return r.query.Manga(ctx, struct{ ID int32 }{ID: r.chapter.MangaID})
}

func newChapterResolver(m *model.Chapter, q *Query) *chapterResolver {
	return &chapterResolver{chapter: m, query: q}
}
