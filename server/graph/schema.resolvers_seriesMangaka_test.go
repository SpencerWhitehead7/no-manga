package graph

import (
	"testing"

	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"
)

// testing the individual fields nested in manga because SeriesMangaka is not a top level query
// it exists only as a child of manga; manga tests the relations, this tests individual fields
func TestSeriesMangaka(t *testing.T) {
	tests := []struct {
		des    string
		setup  func()
		query  string
		expect string
	}{
		{
			des: "resolves id",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mka := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
				testhelpers.MangaToMangaka(t, db, m, mka, "author_artist")
			},
			query:  `{"query":"{manga(ID: 1) {mangakaList {id}}}"}`,
			expect: `{"data":{"manga":{"mangakaList":[{"id":1}]}}}`,
		},
		{
			des: "resolves name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mka := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
				testhelpers.MangaToMangaka(t, db, m, mka, "author_artist")
			},
			query:  `{"query":"{manga(ID: 1) {mangakaList {name}}}"}`,
			expect: `{"data":{"manga":{"mangakaList":[{"name":"tName"}]}}}`,
		},
		{
			des: "resolves [] when no otherNames",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mka := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
				testhelpers.MangaToMangaka(t, db, m, mka, "author_artist")
			},
			query:  `{"query":"{manga(ID: 1) {mangakaList {otherNames}}}"}`,
			expect: `{"data":{"manga":{"mangakaList":[{"otherNames":[]}]}}}`,
		},
		{
			des: "resolves otherNames",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mka := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{OtherNames: []string{"alpha"}})
				testhelpers.MangaToMangaka(t, db, m, mka, "author_artist")
			},
			query:  `{"query":"{manga(ID: 1) {mangakaList {otherNames}}}"}`,
			expect: `{"data":{"manga":{"mangakaList":[{"otherNames":["alpha"]}]}}}`,
		},
		{
			des: "resolves description",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mka := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
				testhelpers.MangaToMangaka(t, db, m, mka, "author_artist")
			},
			query:  `{"query":"{manga(ID: 1) {mangakaList {description}}}"}`,
			expect: `{"data":{"manga":{"mangakaList":[{"description":"tDescription"}]}}}`,
		},
		{
			des: "resolves job",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mka := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
				testhelpers.MangaToMangaka(t, db, m, mka, "author_artist")
			},
			query:  `{"query":"{manga(ID: 1) {mangakaList {job}}}"}`,
			expect: `{"data":{"manga":{"mangakaList":[{"job":"author_artist"}]}}}`,
		},
		{
			des: "resolves mangaList sorted by name",
			setup: func() {
				m1 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{Name: "b"})
				m2 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{Name: "a"})
				mka := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
				testhelpers.MangaToMangaka(t, db, m1, mka, "author_artist")
				testhelpers.MangaToMangaka(t, db, m2, mka, "author_artist")
			},
			query: `{"query":"{manga(ID: 1) {mangakaList {mangaList {name}}}}"}`,
			expect: `{"data":{"manga":{"mangakaList":[{"mangaList":[` +
				`{"name":"a"},` +
				`{"name":"b"}` +
				`]}]}}}`,
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
