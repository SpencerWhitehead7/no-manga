// Package main is a server for handling API requests for no-manga.com.
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/SpencerWhitehead7/no-manga/server/graph"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	db, err := pgxpool.Connect(context.Background(), "postgresql://spencer:@localhost:5432/no-manga")
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/gql", graph.GQLHandler(db))
	r.GET("/", graph.PlaygroundHandler())

	r.Run() // listen and serve on 8080
}
