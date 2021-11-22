package loader

import (
	"strconv"

	"github.com/graph-gophers/dataloader"
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
