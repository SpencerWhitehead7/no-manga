package graph

import (
	"testing"

	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"
)

func TestMagazine(t *testing.T) {
	tests := []struct {
		des    string
		setup  func()
		query  string
		expect string
	}{
		{
			des:    "resolves null when no magazine with given magazineID",
			setup:  func() {},
			query:  `{"query":"{magazine(ID: 1) {id}}"}`,
			expect: `{"data":{"magazine":null}}`,
		},
		{
			des: "resolves id",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazine(ID: 1) {id}}"}`,
			expect: `{"data":{"magazine":{"id":1}}}`,
		},
		{
			des: "resolves name",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazine(ID: 1) {name}}"}`,
			expect: `{"data":{"magazine":{"name":"tName"}}}`,
		},
		{
			des: "resolves [] when no otherNames",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazine(ID: 1) {otherNames}}"}`,
			expect: `{"data":{"magazine":{"otherNames":[]}}}`,
		},
		{
			des: "resolves otherNames",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{OtherNames: []string{"alpha"}})
			},
			query:  `{"query":"{magazine(ID: 1) {otherNames}}"}`,
			expect: `{"data":{"magazine":{"otherNames":["alpha"]}}}`,
		},
		{
			des: "resolves description",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazine(ID: 1) {description}}"}`,
			expect: `{"data":{"magazine":{"description":"tDescription"}}}`,
		},
		{
			des: "resolves demo",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazine(ID: 1) {demo}}"}`,
			expect: `{"data":{"magazine":{"demo":"shonen"}}}`,
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
