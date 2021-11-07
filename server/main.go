// Package main is a server for handling API requests for no-manga.com.
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"github.com/SpencerWhitehead7/no-manga/server/resolver"
	"github.com/SpencerWhitehead7/no-manga/server/schema"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/gql", gin.WrapH(&relay.Handler{
		Schema: graphql.MustParseSchema(schema.Schema, resolver.NewQuery()),
	}))

	r.Run() // listen and serve on 8080
}
