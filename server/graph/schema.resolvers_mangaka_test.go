package graph

import (
	"testing"

	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"
)

func TestMangaka(t *testing.T) {
	tests := []struct {
		des    string
		setup  func()
		query  string
		expect string
	}{
		{
			des:    "resolves null when no mangaka with given mangakaID",
			setup:  func() {},
			query:  `{"query":"{mangaka(ID: 1) {id}}"}`,
			expect: `{"data":{"mangaka":null}}`,
		},
		{
			des: "resolves id",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{mangaka(ID: 1) {id}}"}`,
			expect: `{"data":{"mangaka":{"id":1}}}`,
		},
		{
			des: "resolves name",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{mangaka(ID: 1) {name}}"}`,
			expect: `{"data":{"mangaka":{"name":"tName"}}}`,
		},
		{
			des: "resolves [] when no otherNames",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{mangaka(ID: 1) {otherNames}}"}`,
			expect: `{"data":{"mangaka":{"otherNames":[]}}}`,
		},
		{
			des: "resolves otherNames",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{OtherNames: []string{"alpha"}})
			},
			query:  `{"query":"{mangaka(ID: 1) {otherNames}}"}`,
			expect: `{"data":{"mangaka":{"otherNames":["alpha"]}}}`,
		},
		{
			des: "resolves description",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{mangaka(ID: 1) {description}}"}`,
			expect: `{"data":{"mangaka":{"description":"tDescription"}}}`,
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

func TestMangakaList(t *testing.T) {
	tests := []struct {
		des    string
		setup  func()
		query  string
		expect string
	}{
		{
			des:    "resolves [] when no mangaka",
			setup:  func() {},
			query:  `{"query":"{mangakaList {id}}"}`,
			expect: `{"data":{"mangakaList":[]}}`,
		},
		{
			des: "resolves id",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{mangakaList {id}}"}`,
			expect: `{"data":{"mangakaList":[{"id":1}]}}`,
		},
		{
			des: "resolves name",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{mangakaList {name}}"}`,
			expect: `{"data":{"mangakaList":[{"name":"tName"}]}}`,
		},
		{
			des: "resolves [] when no otherNames",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{mangakaList {otherNames}}"}`,
			expect: `{"data":{"mangakaList":[{"otherNames":[]}]}}`,
		},
		{
			des: "resolves otherNames",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{OtherNames: []string{"alpha"}})
			},
			query:  `{"query":"{mangakaList {otherNames}}"}`,
			expect: `{"data":{"mangakaList":[{"otherNames":["alpha"]}]}}`,
		},
		{
			des: "resolves description",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{mangakaList {description}}"}`,
			expect: `{"data":{"mangakaList":[{"description":"tDescription"}]}}`,
		},
		{
			des: "resolves mangakaList sorted by name",
			setup: func() {
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{Name: "b"})
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{Name: "a"})
			},
			query: `{"query":"{mangakaList {name}}"}`,
			expect: `{"data":{"mangakaList":[` +
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
