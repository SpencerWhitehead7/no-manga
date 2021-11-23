package resolver

import (
	"context"
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type seriesMangakaResolver struct {
	seriesMangaka *model.SeriesMangaka
	query         *Query
}

func (r *seriesMangakaResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.seriesMangaka.ID)))
}
func (r *seriesMangakaResolver) Name() string {
	return r.seriesMangaka.Name
}
func (r *seriesMangakaResolver) OtherNames() []string {
	return r.seriesMangaka.OtherNames
}
func (r *seriesMangakaResolver) Description() string {
	return r.seriesMangaka.Description
}
func (r *seriesMangakaResolver) Job() string {
	return r.seriesMangaka.Job
}
func (r *seriesMangakaResolver) MangaList(ctx context.Context) ([]*mangaResolver, error) {
	return r.query.mangaMListToRList(r.query.loader.MangaListBySeriesMangaka(ctx, r.seriesMangaka))
}
