package main

import (
	"flag"
	"fmt"
	"github.com/mpetavy/common"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

var (
	useGo = flag.Bool("usego", true, "Use GO functionality")
)

func init() {
	common.Init(false, "1.0.0", "", "", "2022", "Btree tryout", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, nil, run, 0)
}

type Entries[T constraints.Ordered] []T

func (entries *Entries[T]) BinarySearch(value T) (int, bool) {
	if *useGo {
		return slices.BinarySearch(*entries, value)
	}

	low := 0
	high := len(*entries) - 1

	for low <= high {
		median := (low + high) / 2

		if (*entries)[median] < value {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(*entries) || (*entries)[low] != value {
		return low, false
	}

	return low, true
}

func (entries *Entries[T]) Insert(value T) {
	index, _ := entries.BinarySearch(value)

	*entries = append(*entries, value)
	copy((*entries)[index+1:], (*entries)[index:])
	(*entries)[index] = value
}

func (entries *Entries[T]) Values() []T {
	slice := make([]T, 0)

	for _, entry := range *entries {
		slice = append(slice, entry)
	}

	return slice
}

type Page[T constraints.Ordered] struct {
	Count   int
	Entries Entries[T]
}

func run() error {
	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
