package main

import (
	"fmt"
	"github.com/mpetavy/common"
	"testing"
)

type Inter struct {
	Value int
}

func (inter Inter) Correlate(value Inter) int {
	if inter.Value < value.Value {
		return -1
	}
	if inter.Value > value.Value {
		return 1
	}

	return 0
}

func TestEntries_Insert(t *testing.T) {
	count := 100000

	values := make([]int, 0)
	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	items := Items[Inter]{}
	slice := make([]int, 0)

	for i := 0; i < count; i++ {
		index := common.Rnd(len(values))

		value := Inter{
			Value: values[index],
		}
		values = append(values[:index], values[index+1:]...)

		items.Insert(value)

		slice = append(slice, value.Value)
	}

	//sort.SliceStable(slice, func(i int, j int) bool {
	//	return slice[i] < slice[j]
	//
	//})
	//
	//assert.Equal(t, slice, items.Values())

	fmt.Printf("%+v\n", items.Values())
}
