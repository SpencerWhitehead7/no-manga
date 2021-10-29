package graph

import (
	"github.com/SpencerWhitehead7/no-manga/server/graph/generated"
	"github.com/SpencerWhitehead7/no-manga/server/graph/repository"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

// if necessary, https://gqlgen.com/recipes/gin/ also has a way to access gin context in resolvers

// GQLHandler returns a function that resolves gql requests.
func GQLHandler(db *pgxpool.Pool) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &Resolver{
					ChapterRepo:       repository.ChapterFactory(db),
					MagazineRepo:      repository.MagazineFactory(db),
					MangaRepo:         repository.MangaFactory(db),
					MangakaRepo:       repository.MangakaFactory(db),
					SeriesMangakaRepo: repository.SeriesMangakaFactory(db),
				},
			},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler returns a function that exposes the gql playground.
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQueryLanguage", "/gql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
