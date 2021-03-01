package graph

import (
	"testing"
	"time"

	"github.com/SpencerWhitehead7/no-manga/server/testhelpers"
)

func TestChapterNoMangaID(t *testing.T) {
	m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
	testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})

	w := testhelpers.CallGQL(r, `{"query":"{chapter(mangaID: 2, chapterNum: 1) {id}}"}`)

	expect := `{"data":{"chapter":null}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}

func TestChapterNoChapterNum(t *testing.T) {
	m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
	testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})

	w := testhelpers.CallGQL(r, `{"query":"{chapter(mangaID: 1, chapterNum: 2) {id}}"}`)

	expect := `{"data":{"chapter":null}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}

func TestChapterNoMangaIdAndNoChapterNum(t *testing.T) {
	m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
	testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})

	w := testhelpers.CallGQL(r, `{"query":"{chapter(mangaID: 2, chapterNum: 2) {id}}"}`)

	expect := `{"data":{"chapter":null}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}

func TestChapterNullFields(t *testing.T) {
	m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
	c := testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m, ChapterNum: 2})

	w := testhelpers.CallGQL(r, `{"query":"{chapter(mangaID: 1, chapterNum: 2) {id chapterNum name pageCount updatedAt}}"}`)

	expect := `{"data":{"chapter":` +
		`{"id":"1__2","chapterNum":2,"name":null,"pageCount":1,"updatedAt":"` + c.UpdatedAt.Format(time.RFC3339) +
		`"}}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}

func TestChapter(t *testing.T) {
	m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
	c := testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m, Name: "tName"})

	w := testhelpers.CallGQL(r, `{"query":"{chapter(mangaID: 1, chapterNum: 1) {id chapterNum name pageCount updatedAt}}"}`)

	expect := `{"data":{"chapter":` +
		`{"id":"1__1","chapterNum":1,"name":"tName","pageCount":1,"updatedAt":"` + c.UpdatedAt.Format(time.RFC3339) +
		`"}}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}

func TestChapterManga(t *testing.T) {
	m := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
	testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m})

	w := testhelpers.CallGQL(r, `{"query":"{chapter(mangaID: 1, chapterNum: 1) {id manga {id}}}"}`)

	expect := `{"data":{"chapter":{"id":"1__1","manga":{"id":1}}}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}

func TestChapterListNoChapters(t *testing.T) {
	w := testhelpers.CallGQL(r, `{"query":"{chapterList {id}}"}`)

	expect := `{"data":{"chapterList":[]}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}
}

func TestChapterList(t *testing.T) {
	m1 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
	m2 := testhelpers.MangaFactory(t, db, testhelpers.MangaStub{})
	c1 := testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, Name: "tName", ChapterNum: 1})
	c2 := testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m1, ChapterNum: 2})
	c3 := testhelpers.ChapterFactory(t, db, testhelpers.ChapterStub{Manga: m2, ChapterNum: 1})

	w := testhelpers.CallGQL(r, `{"query":"{chapterList {id chapterNum name pageCount updatedAt}}"}`)

	expect := `{"data":{"chapterList":[` +
		`{"id":"1__1","chapterNum":1,"name":"tName","pageCount":1,"updatedAt":"` + c1.UpdatedAt.Format(time.RFC3339) + `"},` +
		`{"id":"1__2","chapterNum":2,"name":null,"pageCount":1,"updatedAt":"` + c2.UpdatedAt.Format(time.RFC3339) + `"},` +
		`{"id":"2__1","chapterNum":1,"name":null,"pageCount":1,"updatedAt":"` + c3.UpdatedAt.Format(time.RFC3339) + `"}` +
		`]}}`
	actual := w.Body.String()
	if expect != actual {
		t.Errorf("\nQueryResult\n  expect: %v\n  actual: %v", expect, actual)
	}

	testhelpers.ClearDB(t, db)
}
