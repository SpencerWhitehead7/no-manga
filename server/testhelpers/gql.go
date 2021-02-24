package testhelpers

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

// CallGQL calls / with the provided query and necessary headers for gql and returns the response.
func CallGQL(r *gin.Engine, query string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(query))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	return w
}
