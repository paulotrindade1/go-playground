package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func (t *testing.T) {
		numbers := []int{2, 4, 8, 16, 32}

		got := Sum(numbers)
		want := 62

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{7, 5}, []int{1, 1})
	want := []int{12, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

