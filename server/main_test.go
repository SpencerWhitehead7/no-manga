package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	db := testhelpers.GetDbpool()
	defer db.Close()

	r = getRouter(db)

	os.Exit(m.Run())
}

func callRoute(method string, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var req *http.Request
	var _ error
	if path == "/gql" {
		req, _ = http.NewRequest(method, path, strings.NewReader(`{}`))
		req.Header.Add("Content-Type", "application/json")
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}
	r.ServeHTTP(w, req)

	return w
}

func TestPing(t *testing.T) {
	t.Parallel()
	w := callRoute(http.MethodGet, "/ping")

	expectCode := http.StatusOK
	actualCode := w.Code
	if expectCode != actualCode {
		t.Errorf("\nCode\n  expect: %v\n  actual: %v", expectCode, actualCode)
	}

	expectBody := `{"message":"pong"}`
	actualBody := w.Body.String()
	if expectBody != actualBody {
		t.Errorf("\nBody\n  expect: %v\n  actual: %v", expectBody, actualBody)
	}
}

func TestPlayground(t *testing.T) {
	t.Parallel()
	w := callRoute(http.MethodGet, "/")

	expect := http.StatusOK
	actual := w.Code
	if expect != actual {
		t.Errorf("\nCode\n  expect: %v\n  actual: %v", expect, actual)
	}
}

func TestGQL(t *testing.T) {
	t.Parallel()
	w := callRoute(http.MethodPost, "/gql")

	expect := http.StatusOK
	actual := w.Code
	if expect != actual {
		t.Errorf("\nCode\n  expect: %v\n  actual: %v", expect, actual)
	}
}
