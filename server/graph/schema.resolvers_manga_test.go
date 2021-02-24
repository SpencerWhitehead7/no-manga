package graph

import (
	"testing"

	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"
)

func TestManga(t *testing.T) {
	tests := []struct {
		des    string
		setup  func()
		query  string
		expect string
	}{
		{
			des:    "resolves null when no manga with given mangaID",
			setup:  func() {},
			query:  `{"query":"{manga(ID: 1) {id}}"}`,
			expect: `{"data":{"manga":null}}`,
		},
		{
			des: "resolves id",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{manga(ID: 1) {id}}"}`,
			expect: `{"data":{"manga":{"id":1}}}`,
		},
		{
			des: "resolves name",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{manga(ID: 1) {name}}"}`,
			expect: `{"data":{"manga":{"name":"tName"}}}`,
		},
		{
			des: "resolves [] when no otherNames",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{manga(ID: 1) {otherNames}}"}`,
			expect: `{"data":{"manga":{"otherNames":[]}}}`,
		},
		{
			des: "resolves otherNames",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{OtherNames: []string{"alpha"}})
			},
			query:  `{"query":"{manga(ID: 1) {otherNames}}"}`,
			expect: `{"data":{"manga":{"otherNames":["alpha"]}}}`,
		},
		{
			des: "resolves description",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{manga(ID: 1) {description}}"}`,
			expect: `{"data":{"manga":{"description":"tDescription"}}}`,
		},
		{
			des: "resolves demo",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{manga(ID: 1) {demo}}"}`,
			expect: `{"data":{"manga":{"demo":"shonen"}}}`,
		},
		{
			des: "resolves startDate",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{manga(ID: 1) {startDate}}"}`,
			expect: `{"data":{"manga":{"startDate":"2000-01-01T00:00:00Z"}}}`,
		},
		{
			des: "resolves null when no endDate",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{manga(ID: 1) {endDate}}"}`,
			expect: `{"data":{"manga":{"endDate":null}}}`,
		},
		{
			des: "resolves endDate",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{EndDate: "2020-01-01"})
			},
			query:  `{"query":"{manga(ID: 1) {endDate}}"}`,
			expect: `{"data":{"manga":{"endDate":"2020-01-01T00:00:00Z"}}}`,
		},
	}

	for _, test := range tests {
		test.setup()
		w := testhelpers.CallGQL(r, test.query)
		actual := w.Body.String()
		if test.expect != actual {
			t.Errorf("\n%v\n  expect: %v\n  actual: %v", test.des, test.expect, actual)
		}
		testhelpers.ClearDB(t, db)
	}
}

func TestMangaList(t *testing.T) {
	tests := []struct {
		des    string
		setup  func()
		query  string
		expect string
	}{
		{
			des:    "resolves [] when no manga",
			setup:  func() {},
			query:  `{"query":"{mangaList {id}}"}`,
			expect: `{"data":{"mangaList":[]}}`,
		},
		{
			des: "resolves id",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{mangaList {id}}"}`,
			expect: `{"data":{"mangaList":[{"id":1}]}}`,
		},
		{
			des: "resolves name",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{mangaList {name}}"}`,
			expect: `{"data":{"mangaList":[{"name":"tName"}]}}`,
		},
		{
			des: "resolves [] when no otherNames",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{mangaList {otherNames}}"}`,
			expect: `{"data":{"mangaList":[{"otherNames":[]}]}}`,
		},
		{
			des: "resolves otherNames",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{OtherNames: []string{"alpha"}})
			},
			query:  `{"query":"{mangaList {otherNames}}"}`,
			expect: `{"data":{"mangaList":[{"otherNames":["alpha"]}]}}`,
		},
		{
			des: "resolves description",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{mangaList {description}}"}`,
			expect: `{"data":{"mangaList":[{"description":"tDescription"}]}}`,
		},
		{
			des: "resolves demo",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{mangaList {demo}}"}`,
			expect: `{"data":{"mangaList":[{"demo":"shonen"}]}}`,
		},
		{
			des: "resolves startDate",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{mangaList {startDate}}"}`,
			expect: `{"data":{"mangaList":[{"startDate":"2000-01-01T00:00:00Z"}]}}`,
		},
		{
			des: "resolves null when no endDate",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{mangaList {endDate}}"}`,
			expect: `{"data":{"mangaList":[{"endDate":null}]}}`,
		},
		{
			des: "resolves endDate",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{EndDate: "2020-01-01"})
			},
			query:  `{"query":"{mangaList {endDate}}"}`,
			expect: `{"data":{"mangaList":[{"endDate":"2020-01-01T00:00:00Z"}]}}`,
		},
		{
			des: "resolves mangaList sorted by name",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{Name: "b"})
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{Name: "a"})
			},
			query: `{"query":"{mangaList {name}}"}`,
			expect: `{"data":{"mangaList":[` +
				`{"name":"a"},` +
				`{"name":"b"}` +
				`]}}`,
		},
	}

	for _, test := range tests {
		test.setup()
		w := testhelpers.CallGQL(r, test.query)
		actual := w.Body.String()
		if test.expect != actual {
			t.Errorf("\n%v\n  expect: %v\n  actual: %v", test.des, test.expect, actual)
		}
		testhelpers.ClearDB(t, db)
	}
}