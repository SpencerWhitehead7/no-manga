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
	return r.query.loader.Genres(ctx, r.manga)
}
func (r *mangaResolver) ChapterCount(ctx context.Context) (int32, error) {
	return r.query.chapterRepository.GetCount(ctx, r.manga)
}
func (r *mangaResolver) ChapterList(ctx context.Context) ([]*chapterResolver, error) {
	return r.query.chapterMListToRList(r.query.loader.ChapterListByManga(ctx, r.manga))
}
func (r *mangaResolver) MangakaList(ctx context.Context) ([]*seriesMangakaResolver, error) {
	return r.query.seriesMangakaList(ctx, r.manga)
}
func (r *mangaResolver) MagazineList(ctx context.Context) ([]*magazineResolver, error) {
	return r.query.magazineMListToRList(r.query.magazineRepository.GetByManga(ctx, r.manga))
}
