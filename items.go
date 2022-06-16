package main

type Correlater[T any] interface {
	Correlate(v T) int
}

type Items[T Correlater[T]] struct {
	values []T
}

func (items *Items[T]) BinarySearch(value T) (int, bool) {
	return BinarySearch(items.values, value)
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
	return items.values
}
