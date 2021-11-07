// Package main is a server for handling API requests for no-manga.com.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"

	"github.com/SpencerWhitehead7/no-manga/server/playground"
	"github.com/SpencerWhitehead7/no-manga/server/resolver"
	"github.com/SpencerWhitehead7/no-manga/server/schema"
)

func main() {
	env := os.Getenv("ENV")

	var err error
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		err = godotenv.Load(".env.prod")
	} else {
		err = godotenv.Load(".env.dev")
	}
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("unable to connect to database:", err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalln("database startup ping failed:", err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/dbping", func(c *gin.Context) {
		err := db.Ping(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("database ping failed: %v", err)})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "dbpong"})
		}
	})

	r.POST("/gql", gin.WrapH(&relay.Handler{
		Schema: graphql.MustParseSchema(schema.Schema, resolver.NewQuery()),
	}))
	r.GET("/", playground.NewHandler("no-manga GQL Playground", "/gql"))

	r.Run() // listen and serve on 8080
}
