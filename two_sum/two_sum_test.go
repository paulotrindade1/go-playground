package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("Not found two sum", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		got := TwoSum(input, 20)

		if got != nil {
			t.Errorf("got %q want %q", got, []int{})
		}
	})

	t.Run("Found two sum", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		got := TwoSum(input, 9)

		expected := []int{3, 4}

		assert.Equal(t, expected, got)
	})
}
