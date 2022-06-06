package main

type Page[T Correlater[T]] struct {
	Count int
	Items Items[T]
}
