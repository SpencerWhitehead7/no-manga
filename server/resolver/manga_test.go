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
		{
			des: "resolves [] when no associated genres",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.GenreFactory(t, db, testhelpers.GenreStub{})
			},
			query:  `{"query":"{manga(ID: 1) {genres}}"}`,
			expect: `{"data":{"manga":{"genres":[]}}}`,
		},
		{
			des: "resolves genres sorted by name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				g1 := testhelpers.GenreFactory(t, db, testhelpers.GenreStub{Name: "b"})
				g2 := testhelpers.GenreFactory(t, db, testhelpers.GenreStub{Name: "a"})
				testhelpers.MangaToGenres(t, db, m, []testhelpers.GenreRow{g1, g2})
			},
			query:  `{"query":"{manga(ID: 1) {genres}}"}`,
			expect: `{"data":{"manga":{"genres":["a","b"]}}}`,
		},
		{
			des: "resolves [] when no associated chapters",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				m2 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m2, ChapterNum: 1})
			},
			query:  `{"query":"{manga(ID: 1) {chapterList {id}}}"}`,
			expect: `{"data":{"manga":{"chapterList":[]}}}`,
		},
		{
			des: "resolves chapters sorted by chapterNum",
			setup: func() {
				m1 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 2})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 1})
			},
			query: `{"query":"{manga(ID: 1) {chapterList {id}}}"}`,
			expect: `{"data":{"manga":{"chapterList":[` +
				`{"id":"1__1"},` +
				`{"id":"1__2"}` +
				`]}}}`,
		},
		{
			des: "resolves 0 when no associated chapters",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				m2 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m2, ChapterNum: 1})
			},
			query:  `{"query":"{manga(ID: 1) {chapterCount}}"}`,
			expect: `{"data":{"manga":{"chapterCount":0}}}`,
		},
		{
			des: "resolves chapterCount",
			setup: func() {
				m1 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 1})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 2})
			},
			query:  `{"query":"{manga(ID: 1) {chapterCount}}"}`,
			expect: `{"data":{"manga":{"chapterCount":2}}}`,
		},
		{
			des: "resolves [] when no associated mangaka",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{manga(ID: 1) {mangakaList {id}}}"}`,
			expect: `{"data":{"manga":{"mangakaList":[]}}}`,
		},
		{
			des: "resolves mangaka sorted by name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mka1 := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{Name: "b"})
				mka2 := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{Name: "a"})
				testhelpers.MangaToMangaka(t, db, m, mka1, "author")
				testhelpers.MangaToMangaka(t, db, m, mka2, "artist")
			},
			query: `{"query":"{manga(ID: 1) {mangakaList {name}}}"}`,
			expect: `{"data":{"manga":{"mangakaList":[` +
				`{"name":"a"},` +
				`{"name":"b"}` +
				`]}}}`,
		},
		{
			des: "resolves [] when no associated magazines",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{manga(ID: 1) {magazineList {id}}}"}`,
			expect: `{"data":{"manga":{"magazineList":[]}}}`,
		},
		{
			des: "resolves magazines sorted by name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mag1 := testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{Name: "b"})
				mag2 := testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{Name: "a"})
				testhelpers.MangaToMagazine(t, db, m, mag1)
				testhelpers.MangaToMagazine(t, db, m, mag2)
			},
			query: `{"query":"{manga(ID: 1) {magazineList {name}}}"}`,
			expect: `{"data":{"manga":{"magazineList":[` +
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
			expect: `{"data":{"mangaList":[{"id":"1"}]}}`,
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
			des: "resolves [] when no associated genres",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.GenreFactory(t, db, testhelpers.GenreStub{})
			},
			query:  `{"query":"{mangaList {genres}}"}`,
			expect: `{"data":{"mangaList":[{"genres":[]}]}}`,
		},
		{
			des: "resolves genres sorted by name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				g1 := testhelpers.GenreFactory(t, db, testhelpers.GenreStub{Name: "b"})
				g2 := testhelpers.GenreFactory(t, db, testhelpers.GenreStub{Name: "a"})
				testhelpers.MangaToGenres(t, db, m, []testhelpers.GenreRow{g1, g2})
			},
			query:  `{"query":"{mangaList {genres}}"}`,
			expect: `{"data":{"mangaList":[{"genres":["a","b"]}]}}`,
		},
		{
			des: "resolves [] when no associated chapters",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				m2 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m2, ChapterNum: 1})
			},
			query:  `{"query":"{mangaList {chapterList {id}}}"}`,
			expect: `{"data":{"mangaList":[{"chapterList":[]},{"chapterList":[{"id":"2__1"}]}]}}`,
		},
		{
			des: "resolves chapters sorted by chapterNum",
			setup: func() {
				m1 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 2})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 1})
			},
			query: `{"query":"{mangaList {chapterList {id}}}"}`,
			expect: `{"data":{"mangaList":[{"chapterList":[` +
				`{"id":"1__1"},` +
				`{"id":"1__2"}` +
				`]}]}}`,
		},
		{
			des: "resolves 0 when no associated chapters",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				m2 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m2, ChapterNum: 1})
			},
			query:  `{"query":"{mangaList {chapterCount}}"}`,
			expect: `{"data":{"mangaList":[{"chapterCount":0},{"chapterCount":1}]}}`,
		},
		{
			des: "resolves chapterCount",
			setup: func() {
				m1 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 1})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 2})
			},
			query:  `{"query":"{mangaList {chapterCount}}"}`,
			expect: `{"data":{"mangaList":[{"chapterCount":2}]}}`,
		},
		{
			des: "resolves [] when no associated mangaka",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{})
			},
			query:  `{"query":"{mangaList {mangakaList {id}}}"}`,
			expect: `{"data":{"mangaList":[{"mangakaList":[]}]}}`,
		},
		{
			des: "resolves mangaka sorted by name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mka1 := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{Name: "b"})
				mka2 := testhelpers.MangakaFactory(t, db, testhelpers.MangakaStub{Name: "a"})
				testhelpers.MangaToMangaka(t, db, m, mka1, "author")
				testhelpers.MangaToMangaka(t, db, m, mka2, "artist")
			},
			query: `{"query":"{mangaList {mangakaList {name}}}"}`,
			expect: `{"data":{"mangaList":[{"mangakaList":[` +
				`{"name":"a"},` +
				`{"name":"b"}` +
				`]}]}}`,
		},
		{
			des: "resolves [] when no associated magazines",
			setup: func() {
				testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{})
			},
			query:  `{"query":"{mangaList {magazineList {id}}}"}`,
			expect: `{"data":{"mangaList":[{"magazineList":[]}]}}`,
		},
		{
			des: "resolves magazines sorted by name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				mag1 := testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{Name: "b"})
				mag2 := testhelpers.MagazineFactory(t, db, testhelpers.MagazineStub{Name: "a"})
				testhelpers.MangaToMagazine(t, db, m, mag1)
				testhelpers.MangaToMagazine(t, db, m, mag2)
			},
			query: `{"query":"{mangaList {magazineList {name}}}"}`,
			expect: `{"data":{"mangaList":[{"magazineList":[` +
				`{"name":"a"},` +
				`{"name":"b"}` +
				`]}]}}`,
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
