package resolver

import (
	"context"
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type magazineResolver struct {
	magazine *model.Magazine
	query    *Query
}

func (r *magazineResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.magazine.ID)))
}
func (r *magazineResolver) Name() string {
	return r.magazine.Name
}
func (r *magazineResolver) OtherNames() []string {
	return r.magazine.OtherNames
}
func (r *magazineResolver) Description() string {
	return r.magazine.Description
}
func (r *magazineResolver) Demo() string {
	return r.magazine.Demo
}
func (r *magazineResolver) MangaList(ctx context.Context) ([]*mangaResolver, error) {
	return r.query.mangaMListToRList(r.query.mangaRepository.GetByMagazine(ctx, r.magazine))
}
