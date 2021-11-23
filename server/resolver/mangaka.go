package resolver

import (
	"context"
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type mangakaResolver struct {
	mangaka *model.Mangaka
	query   *Query
}

func (r *mangakaResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.mangaka.ID)))
}
func (r *mangakaResolver) Name() string {
	return r.mangaka.Name
}
func (r *mangakaResolver) OtherNames() []string {
	return r.mangaka.OtherNames
}
func (r *mangakaResolver) Description() string {
	return r.mangaka.Description
}
func (r *mangakaResolver) MangaList(ctx context.Context) ([]*mangaResolver, error) {
	return r.query.mangaMListToRList(r.query.loader.MangaListByMangaka(ctx, r.mangaka))
}
