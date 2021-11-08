package resolver

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
			expect: `{"data":{"manga":{"id":"1"}}}`,
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
