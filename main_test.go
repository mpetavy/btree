package main

import (
	"fmt"
	"github.com/mpetavy/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	count = 100000
)

type IntValue struct {
	Value int
}

func (intValue IntValue) Correlate(value IntValue) int {
	if intValue.Value < value.Value {
		return -1
	}
	if intValue.Value > value.Value {
		return 1
	}

	return 0
}

type StringValue struct {
	Value string
}

func (stringValue StringValue) Correlate(value StringValue) int {
	if stringValue.Value < value.Value {
		return -1
	}
	if stringValue.Value > value.Value {
		return 1
	}

	return 0
}

func TestIntValues_Insert(t *testing.T) {
	items := Items[IntValue]{}

	for i := 0; i < count; i++ {
		items.Insert(IntValue{Value: common.Rnd(common.MaxInt)})
	}

	lastValue := IntValue{Value: 0}
	for _, intValue := range items.Values() {
		fmt.Printf("%+v\n", intValue.Value)

		assert.Truef(t, lastValue.Value <= intValue.Value, "must be smaller")
	}
}

func TestStringValues_Insert(t *testing.T) {
	items := Items[StringValue]{}

	for i := 0; i < count; i++ {
		str, err := common.RndString(10)

		if err != nil {
			t.Error(err)
		}

		items.Insert(StringValue{Value: str})
	}

	lastValue := StringValue{Value: ""}
	for _, stringValue := range items.Values() {
		fmt.Printf("%+v\n", stringValue.Value)

		assert.Truef(t, lastValue.Value <= stringValue.Value, "must be smaller")
	}
}
