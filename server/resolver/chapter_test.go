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

func TestChapterList(t *testing.T) {
	var testC testhelpers.ChapterRow

	tests := []struct {
		des    string
		setup  func()
		query  string
		expect func() string
	}{
		{
			des:   "returns [] when no chapters",
			setup: func() {},
			query: `{"query":"{chapterList {id}}"}`,
			expect: func() string {
				return `{"data":{"chapterList":[]}}`
			},
		},
		{
			des: "resolves id",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m, ChapterNum: 2})
			},
			query: `{"query":"{chapterList {id}}"}`,
			expect: func() string {
				return `{"data":{"chapterList":[{"id":"1__2"}]}}`
			},
		},
		{
			des: "resolves null when no name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})
			},
			query: `{"query":"{chapterList {name}}"}`,
			expect: func() string {
				return `{"data":{"chapterList":[{"name":null}]}}`
			},
		},
		{
			des: "resolves name",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m, Name: "tName"})
			},
			query: `{"query":"{chapterList {name}}"}`,
			expect: func() string {
				return `{"data":{"chapterList":[{"name":"tName"}]}}`
			},
		},
		{
			des: "resolves pageCount",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})
			},
			query: `{"query":"{chapterList {pageCount}}"}`,
			expect: func() string {
				return `{"data":{"chapterList":[{"pageCount":1}]}}`
			},
		},
		{
			des: "resolves updatedAt",
			setup: func() {
				m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testC = testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})
			},
			query: `{"query":"{chapterList {updatedAt}}"}`,
			expect: func() string {
				return `{"data":{"chapterList":[{"updatedAt":"` + testC.UpdatedAt.Format(time.RFC3339Nano) + `"}]}}`
			},
		},
		{
			des: "resolves chapterList sorted by updatedAt",
			setup: func() {
				m1 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				m2 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 2})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m2, ChapterNum: 1})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 1})
				testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m2, ChapterNum: 2})
			},
			query: `{"query":"{chapterList {id}}"}`,
			expect: func() string {
				return `{"data":{"chapterList":[` +
					`{"id":"2__2"},` +
					`{"id":"1__1"},` +
					`{"id":"2__1"},` +
					`{"id":"1__2"}` +
					`]}}`
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
