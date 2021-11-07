package resolver

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type seriesMangakaResolver struct {
	seriesMangaka *model.SeriesMangaka
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
func (r *seriesMangakaResolver) MangaList() []*mangaResolver {
	return make([]*mangaResolver, 0)
}
