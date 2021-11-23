package loader

import (
	"strconv"

	"github.com/graph-gophers/dataloader"

	"github.com/SpencerWhitehead7/no-manga/server/model"
)

type int32Key int32

func (k int32Key) String() string {
	return strconv.FormatInt(int64(k), 10)
}
func (k int32Key) Raw() interface{} {
	return int32(k)
}

func int32KeysToIDs(keys dataloader.Keys) []int32 {
	ids := make([]int32, len(keys))
	for i, k := range keys {
		ids[i] = k.Raw().(int32)
	}
	return ids
}

type chapterKey model.ChapterID

func (k chapterKey) String() string {
	return k.Raw().(model.ChapterID).ID()
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
