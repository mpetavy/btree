package main

import (
	"fmt"
	"github.com/mpetavy/common"
	"github.com/stretchr/testify/assert"
	"strings"
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
	me := strings.ToUpper(stringValue.Value)
	other := strings.ToUpper(value.Value)

	if me < other {
		return -1
	}
	if me > other {
		return 1
	}

	return 0
}

func TestIntValues_Insert(t *testing.T) {
	items := Items[IntValue]{}

	for i := 0; i < count; i++ {
		if i%1000 == 0 {
			fmt.Printf("%d\n", i)
		}

		items.Insert(IntValue{Value: common.Rnd(common.MaxInt)})
	}

	lastValue := IntValue{Value: 0}
	for _, intValue := range items.Values() {
		assert.Truef(t, lastValue.Value <= intValue.Value, "must be smaller")
	}
}

func TestStringValues_Insert(t *testing.T) {
	items := Items[StringValue]{}

	for i := 0; i < count; i++ {
		if i%1000 == 0 {
			fmt.Printf("%d\n", i)
		}

		str, err := common.RndString(10)

		if err != nil {
			t.Error(err)
		}

		items.Insert(StringValue{Value: str})
	}

	lastValue := StringValue{Value: ""}
	for _, stringValue := range items.Values() {
		assert.Truef(t, lastValue.Value <= stringValue.Value, "must be smaller")
	}
}
