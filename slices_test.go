package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func makeIntSlice(low int, high int) []int {
	slice := make([]int, 0)
	for i := low; i < high; i++ {
		slice = append(slice, i)
	}

	return slice
}

func TestAdd(t *testing.T) {
	type test struct {
		slice    []int
		values   []int
		expected []int
	}
	tests := []test{}

	for i := 0; i < 10; i++ {
		tests = append(tests,
			test{
				slice:    []int{},
				values:   makeIntSlice(i, i+1),
				expected: makeIntSlice(i, i+1),
			})
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add %+v", test.values), func(t *testing.T) {
			Add(&test.slice, test.expected...)

			require.Equal(t, test.expected, test.slice)
		})
	}
}

func TestInsert(t *testing.T) {
	type test struct {
		slice    []int
		index    int
		values   []int
		expected []int
	}
	tests := []test{
		test{
			slice:    []int{},
			index:    0,
			values:   []int{1, 2},
			expected: []int{1, 2},
		},
		test{
			slice:    []int{1},
			index:    1,
			values:   []int{2, 3},
			expected: []int{1, 2, 3},
		},
		test{
			slice:    []int{1},
			index:    0,
			values:   []int{2, 3},
			expected: []int{2, 3, 1},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Insert %+v at %d", test.values, i), func(t *testing.T) {
			Insert(&test.slice, test.index, test.values...)

			require.Equal(t, test.expected, test.slice)
		})
	}
}

//func TestInsert(t *testing.T) {
//	type args struct {
//		slice  *[]T
//		index  int
//		values []T
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			Insert(tt.args.slice, tt.args.index, tt.args.values...)
//		})
//	}
//}
//
//func TestPop(t *testing.T) {
//	type args struct {
//		slice *[]T
//	}
//	tests := []struct {
//		name string
//		args args
//		want T
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			require.Equalf(t, tt.want, Pop(tt.args.slice), "Pop(%v)", tt.args.slice)
//		})
//	}
//}
//
//func TestPopFront(t *testing.T) {
//	type args struct {
//		slice *[]T
//	}
//	tests := []struct {
//		name string
//		args args
//		want T
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			require.Equalf(t, tt.want, PopFront(tt.args.slice), "PopFront(%v)", tt.args.slice)
//		})
//	}
//}
//
//func TestPush(t *testing.T) {
//	type args struct {
//		slice *[]T
//		value T
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			Push(tt.args.slice, tt.args.value)
//		})
//	}
//}
//
//func TestPushFront(t *testing.T) {
//	type args struct {
//		slice  *[]T
//		values []T
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			PushFront(tt.args.slice, tt.args.values...)
//		})
//	}
//}
//
//func TestRemove(t *testing.T) {
//	type args struct {
//		slice *[]T
//		index int
//	}
//	tests := []struct {
//		name string
//		args args
//		want T
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			require.Equalf(t, tt.want, Remove(tt.args.slice, tt.args.index), "Remove(%v, %v)", tt.args.slice, tt.args.index)
//		})
//	}
//}
//
//func TestRemoveRange(t *testing.T) {
//	type args struct {
//		slice *[]T
//		low   int
//		high  int
//	}
//	tests := []struct {
//		name string
//		args args
//		want []T
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			require.Equalf(t, tt.want, RemoveRange(tt.args.slice, tt.args.low, tt.args.high), "RemoveRange(%v, %v, %v)", tt.args.slice, tt.args.low, tt.args.high)
//		})
//	}
//}
