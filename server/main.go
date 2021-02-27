// Package main is a server for handling API requests for no-manga.com.
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/SpencerWhitehead7/no-manga/server/graph"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func getRouter(db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/gql", graph.GQLHandler(db))
	r.GET("/", graph.PlaygroundHandler())

	return r
}

func main() {
	env := os.Getenv("ENV")

	var err error
	if env == "prod" {
		err = godotenv.Load(".env.prod")
	} else {
		err = godotenv.Load(".env.dev")
	}
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	defer db.Close()

	r := getRouter(db)

	r.Run() // listen and serve on 8080
}
