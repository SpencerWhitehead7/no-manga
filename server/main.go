// Package main is a server for handling API requests for no-manga.com.
package main

import (
	"net/http"

	"github.com/SpencerWhitehead7/no-manga/server/graph"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/gql", graph.GQLHandler())
	r.GET("/", graph.PlaygroundHandler())

	r.Run() // listen and serve on 8080
}
