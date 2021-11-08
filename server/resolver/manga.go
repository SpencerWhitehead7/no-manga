package resolver

import (
	"context"
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type mangaResolver struct {
	manga *model.Manga
	query *Query
}

func (r *mangaResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.manga.ID)))
}
func (r *mangaResolver) Name() string {
	return r.manga.Name
}
func (r *mangaResolver) OtherNames() []string {
	return r.manga.OtherNames
}
func (r *mangaResolver) Description() string {
	return r.manga.Description
}
func (r *mangaResolver) Demo() string {
	return r.manga.Demo
}
func (r *mangaResolver) StartDate() graphql.Time {
	return graphql.Time{Time: r.manga.StartDate}
}
func (r *mangaResolver) EndDate() *graphql.Time {
	if r.manga.EndDate != nil {
		return &graphql.Time{Time: *r.manga.EndDate}
	}
	return nil
}
func (r *mangaResolver) Genres(ctx context.Context) ([]string, error) {
	return r.query.mangaRepository.GetGenres(ctx, r.manga)
}
func (r *mangaResolver) ChapterCount() int32 {
	return r.manga.ChapterCount
}
func (r *mangaResolver) ChapterList(ctx context.Context) ([]*chapterResolver, error) {
	return r.query.chapterMListToRList(r.query.chapterRepository.GetByManga(ctx, r.manga))
}
func (r *mangaResolver) MangakaList() []*seriesMangakaResolver {
	return make([]*seriesMangakaResolver, 0)
}
func (r *mangaResolver) MagazineList() []*magazineResolver {
	return make([]*magazineResolver, 0)
}
