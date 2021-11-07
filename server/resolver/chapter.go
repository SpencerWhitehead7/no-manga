package resolver

import (
	"github.com/graph-gophers/graphql-go"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type chapterResolver struct {
	chapter *model.Chapter
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
func (r *chapterResolver) Manga() *mangaResolver {
	return &mangaResolver{}
}
