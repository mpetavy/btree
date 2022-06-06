package main

func Add[T any](slice *[]T, values ...T) {
	*slice = append(*slice, values...)
}

func Insert[T any](slice *[]T, index int, values ...T) {
	*slice = append(*slice, values...)
	copy((*slice)[index+len(values):], (*slice)[index:])
	copy((*slice)[index:], values)
}

func Remove[T any](slice *[]T, index int) T {
	return RemoveRange(slice, index, index+1)[0]
}

func RemoveRange[T any](slice *[]T, low int, high int) []T {
	r := (*slice)[low:high]

	*slice = append((*slice)[:low-1], (*slice)[high:]...)

	return r
}

func Push[T any](slice *[]T, value T) {
	Add(slice, value)
}

func Pop[T any](slice *[]T) T {
	var value T

	value, *slice = (*slice)[len(*slice)-1], (*slice)[:len(*slice)-1]

	return value
}

func PushFront[T any](slice *[]T, values ...T) {
	Insert(slice, 0, values...)
}

func PopFront[T any](slice *[]T) T {
	return Remove(slice, 0)
}

func BinarySearch[T Correlater[T]](entries []T, value T) (int, bool) {
	low := 0
	high := len(entries) - 1

	for low <= high {
		median := (low + high) / 2

		if entries[median].Correlate(value) < 0 {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(entries) || entries[low].Correlate(value) != 0 {
		return low, false
	}

	return low, true
}
