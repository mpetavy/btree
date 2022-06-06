package main

import "golang.org/x/exp/constraints"

type Page[T constraints.Ordered] struct {
	Count int
	Items Items[T]
}
