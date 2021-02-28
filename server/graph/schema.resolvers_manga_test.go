package graph

import (
	"testing"

	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"
)

func TestMangaNoManga(t *testing.T) {
	w := testhelpers.CallGQL(r, `{"query":"{manga(ID: 1) {id}}"}`)

	expect := `{"data":{"manga":null}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}
}

func TestMangaNullFields(t *testing.T) {
	testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})

	w := testhelpers.CallGQL(r, `{"query":"{manga(ID: 1) {id name otherNames description demo startDate endDate}}"}`)

	expect := `{"data":{"manga":` +
		`{"id":1,"name":"tName","otherNames":[],"description":"tDescription","demo":"shonen","startDate":"2000-01-01T00:00:00Z","endDate":null}` +
		`}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}

func TestManga(t *testing.T) {
	testhelpers.MangaFactory(t, db, testhelpers.MangaStub{
		OtherNames: []string{"alpha"}, EndDate: "2020-01-01",
	})

	w := testhelpers.CallGQL(r, `{"query":"{manga(ID: 1) {id name otherNames description demo startDate endDate}}"}`)

	expect := `{"data":{"manga":` +
		`{"id":1,"name":"tName","otherNames":["alpha"],"description":"tDescription","demo":"shonen","startDate":"2000-01-01T00:00:00Z","endDate":"2020-01-01T00:00:00Z"}` +
		`}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}

func TestMangaListNoMangas(t *testing.T) {
	w := testhelpers.CallGQL(r, `{"query":"{mangaList {id}}"}`)

	expect := `{"data":{"mangaList":[]}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}
}

func TestMangaList(t *testing.T) {
	testhelpers.MangaFactory(t, db, testhelpers.MangaStub{
		Name: "b", EndDate: "2020-01-01",
	})
	testhelpers.MangaFactory(t, db, testhelpers.MangaStub{
		Name: "a", OtherNames: []string{"alpha"},
	})

	w := testhelpers.CallGQL(r, `{"query":"{mangaList {id name otherNames description demo startDate endDate}}"}`)

	expect := `{"data":{"mangaList":[` +
		`{"id":2,"name":"a","otherNames":["alpha"],"description":"tDescription","demo":"shonen","startDate":"2000-01-01T00:00:00Z","endDate":null},` +
		`{"id":1,"name":"b","otherNames":[],"description":"tDescription","demo":"shonen","startDate":"2000-01-01T00:00:00Z","endDate":"2020-01-01T00:00:00Z"}` +
		`]}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}
