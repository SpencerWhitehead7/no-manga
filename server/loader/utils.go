package loader

import (
	"strconv"

	"github.com/graph-gophers/dataloader"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

func int32KeysToIDs(keys dataloader.Keys) []int32 {
	ids := make([]int32, len(keys))
	for i, k := range keys {
		ids[i] = k.Raw().(int32)
	}
	return ids
}

type chapterKey model.ChapterID

func (k chapterKey) String() string {
	return "chapter_" + k.Raw().(model.ChapterID).ID()
}
func (k chapterKey) Raw() interface{} {
	return model.ChapterID(k)
}

func chapterKeysToIDs(keys dataloader.Keys) []model.ChapterID {
	ids := make([]model.ChapterID, len(keys))
	for i, k := range keys {
		ids[i] = k.Raw().(model.ChapterID)
	}
	return ids
}

type chapterCountKey int32

func (k chapterCountKey) String() string {
	return "chapterCount_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k chapterCountKey) Raw() interface{} {
	return int32(k)
}

type chapterListByMangaKey int32

func (k chapterListByMangaKey) String() string {
	return "chapterListByManga_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k chapterListByMangaKey) Raw() interface{} {
	return int32(k)
}

type genresKey int32

func (k genresKey) String() string {
	return "genres_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k genresKey) Raw() interface{} {
	return int32(k)
}

type magazineKey int32

func (k magazineKey) String() string {
	return "magazine_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k magazineKey) Raw() interface{} {
	return int32(k)
}

type magazineListByMangaKey int32

func (k magazineListByMangaKey) String() string {
	return "magazineListByManga_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k magazineListByMangaKey) Raw() interface{} {
	return int32(k)
}

type mangaKey int32

func (k mangaKey) String() string {
	return "manga_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k mangaKey) Raw() interface{} {
	return int32(k)
}

type mangaListByMagazineKey int32

func (k mangaListByMagazineKey) String() string {
	return "mangaListByMagazine_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k mangaListByMagazineKey) Raw() interface{} {
	return int32(k)
}

type mangaListByMangakaKey int32

func (k mangaListByMangakaKey) String() string {
	return "mangaListByMangaka_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k mangaListByMangakaKey) Raw() interface{} {
	return int32(k)
}

type mangaListBySeriesMangakaKey int32

func (k mangaListBySeriesMangakaKey) String() string {
	return "mangaListBySeriesMangaka_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k mangaListBySeriesMangakaKey) Raw() interface{} {
	return int32(k)
}

type mangakaKey int32

func (k mangakaKey) String() string {
	return "mangaka_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k mangakaKey) Raw() interface{} {
	return int32(k)
}

type seriesMangakaListKey int32

func (k seriesMangakaListKey) String() string {
	return "seriesMangakaList_" + strconv.FormatInt(int64(k.Raw().(int32)), 10)
}
func (k seriesMangakaListKey) Raw() interface{} {
	return int32(k)
}

func loadBatchError(keys dataloader.Keys, err error) []*dataloader.Result {
	r := &dataloader.Result{Error: err}

	res := make([]*dataloader.Result, len(keys))
	for i := range keys {
		res[i] = r
	}

	return res
}

// func loadBatchSuccess() TODO generics :/

func handleSingleBatch(keys dataloader.Keys, list interface{}, err error) []*dataloader.Result {
	if err != nil {
		return loadBatchError(keys, err)
	}

	return []*dataloader.Result{{Data: list}}
}
