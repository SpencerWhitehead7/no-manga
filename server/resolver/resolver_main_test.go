package resolver

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/SpencerWhitehead7/no-manga/server/schema"
	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"
)

var db *pgxpool.Pool
var r *gin.Engine

func TestMain(m *testing.M) {
	db = testhelpers.GetDbpool()
	defer db.Close()

	r = gin.Default()
	r.POST("/", gin.WrapH(&relay.Handler{
		Schema: graphql.MustParseSchema(schema.Schema, NewQuery(db)),
	}))

	os.Exit(m.Run())
}
