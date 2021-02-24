package graph

import (
	"os"
	"testing"

	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool
var r *gin.Engine

func TestMain(m *testing.M) {
	db = testhelpers.GetDbpool()
	defer db.Close()

	r = gin.Default()
	r.POST("/", GQLHandler(db))

	os.Exit(m.Run())
}
