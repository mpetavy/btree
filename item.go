package main

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Items[T constraints.Ordered] struct {
	values []T
}

func (items *Items[T]) BinarySearch(value T) (int, bool) {
	return slices.BinarySearch(items.values, value)
}

func (items *Items[T]) Insert(value T) {
	index, _ := items.BinarySearch(value)

	Insert(&items.values, index, value)
}

func (items *Items[T]) Remove(value T) {
	index, _ := items.BinarySearch(value)

	Remove(&items.values, index)
}

func (items *Items[T]) Values() []T {
	slice := make([]T, 0, len(items.values))

	for _, entry := range items.values {
		slice = append(slice, entry)
	}

	return slice
}
