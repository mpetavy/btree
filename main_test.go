package main

import (
	"github.com/mpetavy/common"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestEntries_Insert(t *testing.T) {
	count := 100000

	values := make([]int, 0)
	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	entries := Entries[int]{}
	slice := make([]int, 0)

	for i := 0; i < count; i++ {
		index := common.Rnd(len(values))

		value := values[index]
		values = append(values[:index], values[index+1:]...)

		entries.Insert(value)

		slice = append(slice, value)
	}

	sort.SliceStable(slice, func(i int, j int) bool {
		return slice[i] < slice[j]

	})

	assert.Equal(t, slice, entries.Values())
}
