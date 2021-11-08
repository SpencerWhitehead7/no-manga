package resolver

import (
	"testing"
	"time"

	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"
)

func TestChapter(t *testing.T) {
	var testC testhelpers.ChapterRow

	tests := []struct {
		des    string
		setup  func()
		query  string
		expect func() string
	}{
		{
			des: "resolves null when no manga with given mangaID",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})
			},
			query: `{"query":"{chapter(mangaID: 2, chapterNum: 1) {id}}"}`,
			expect: func() string {
				return `{"data":{"chapter":null}}`
			},
		},
		{
			des: "resolves null when no chapter with given ChapterNum for given mangaID",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})
			},
			query: `{"query":"{chapter(mangaID: 1, chapterNum: 2) {id}}"}`,
			expect: func() string {
				return `{"data":{"chapter":null}}`
			},
		},
		{
			des: "resolves null when no manga with given mangaID and no chapter with given chapterNum",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})
			},
			query: `{"query":"{chapter(mangaID: 2, chapterNum: 2) {id}}"}`,
			expect: func() string {
				return `{"data":{"chapter":null}}`
			},
		},
		{
			des: "resolves id",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m, ChapterNum: 2})
			},
			query: `{"query":"{chapter(mangaID: 1, chapterNum: 2) {id}}"}`,
			expect: func() string {
				return `{"data":{"chapter":{"id":"1__2"}}}`
			},
		},
		{
			des: "resolves null when no name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})
			},
			query: `{"query":"{chapter(mangaID: 1, chapterNum: 1) {name}}"}`,
			expect: func() string {
				return `{"data":{"chapter":{"name":null}}}`
			},
		},
		{
			des: "resolves name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m, Name: "tName"})
			},
			query: `{"query":"{chapter(mangaID: 1, chapterNum: 1) {name}}"}`,
			expect: func() string {
				return `{"data":{"chapter":{"name":"tName"}}}`
			},
		},
		{
			des: "resolves pageCount",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})
			},
			query: `{"query":"{chapter(mangaID: 1, chapterNum: 1) {pageCount}}"}`,
			expect: func() string {
				return `{"data":{"chapter":{"pageCount":1}}}`
			},
		},
		{
			des: "resolves updatedAt",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})
			},
			query: `{"query":"{chapter(mangaID: 1, chapterNum: 1) {updatedAt}}"}`,
			expect: func() string {
				return `{"data":{"chapter":{"updatedAt":"` + testC.UpdatedAt.Format(time.RFC3339Nano) + `"}}}`
			},
		},
	}

	for _, test := range tests {
		test.setup()
		w := testhelpers.CallGQL(r, test.query)
		actual := w.Body.String()
		if test.expect() != actual {
			t.Errorf("\n%v\n  expect: %v\n  actual: %v", test.des, test.expect(), actual)
		}
		testhelpers.ClearDB(t, db)
	}
}
