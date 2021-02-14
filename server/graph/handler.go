package graph

import (
	"github.com/SpencerWhitehead7/no-manga/server/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// if necessary, https://gqlgen.com/recipes/gin/ also has a way to access gin context in resolvers

// GQLHandler returns a function that resolves gql requests.
func GQLHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &Resolver{},
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
