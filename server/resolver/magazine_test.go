package resolver

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
			expect: `{"data":{"magazine":{"id":"1"}}}`,
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
		{
			des: "resolves [] when no associated manga",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{magazine(ID: 1) {mangaList {id}}}"}`,
			expect: `{"data":{"magazine":{"mangaList":[]}}}`,
		},
		{
			des: "resolves manga sorted by name",
			setup: func() {
				mag := testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
				m1 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{Name: "b"})
				m2 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{Name: "a"})
				testhelpers.MangaToMagazine(t, db, m1, mag)
				testhelpers.MangaToMagazine(t, db, m2, mag)
			},
			query: `{"query":"{magazine(ID: 1) {mangaList {name}}}"}`,
			expect: `{"data":{"magazine":{"mangaList":[` +
				`{"name":"a"},` +
				`{"name":"b"}` +
				`]}}}`,
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

func TestMagazineList(t *testing.T) {
	tests := []struct {
		des    string
		setup  func()
		query  string
		expect string
	}{
		{
			des:    "resolves [] when no magazines",
			setup:  func() {},
			query:  `{"query":"{magazineList {id}}"}`,
			expect: `{"data":{"magazineList":[]}}`,
		},
		{
			des: "resolves id",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazineList {id}}"}`,
			expect: `{"data":{"magazineList":[{"id":"1"}]}}`,
		},
		{
			des: "resolves name",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazineList {name}}"}`,
			expect: `{"data":{"magazineList":[{"name":"tName"}]}}`,
		},
		{
			des: "resolves [] when no otherNames",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazineList {otherNames}}"}`,
			expect: `{"data":{"magazineList":[{"otherNames":[]}]}}`,
		},
		{
			des: "resolves otherNames",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{OtherNames: []string{"alpha"}})
			},
			query:  `{"query":"{magazineList {otherNames}}"}`,
			expect: `{"data":{"magazineList":[{"otherNames":["alpha"]}]}}`,
		},
		{
			des: "resolves description",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazineList {description}}"}`,
			expect: `{"data":{"magazineList":[{"description":"tDescription"}]}}`,
		},
		{
			des: "resolves demo",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{magazineList {demo}}"}`,
			expect: `{"data":{"magazineList":[{"demo":"shonen"}]}}`,
		},
		{
			des: "resolves [] when no associated manga",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
			},
			query:  `{"query":"{magazineList {mangaList {id}}}"}`,
			expect: `{"data":{"magazineList":[{"mangaList":[]}]}}`,
		},
		{
			des: "resolves manga sorted by name",
			setup: func() {
				mag := testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
				m1 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{Name: "b"})
				m2 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{Name: "a"})
				testhelpers.MangaToMagazine(t, db, m1, mag)
				testhelpers.MangaToMagazine(t, db, m2, mag)
			},
			query: `{"query":"{magazineList {mangaList {name}}}"}`,
			expect: `{"data":{"magazineList":[{"mangaList":[` +
				`{"name":"a"},` +
				`{"name":"b"}` +
				`]}]}}`,
		},
		{
			des: "resolves magazineList sorted by name",
			setup: func() {
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{Name: "b"})
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{Name: "a"})
			},
			query: `{"query":"{magazineList {name}}"}`,
			expect: `{"data":{"magazineList":[` +
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
