package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	assertCorrectResult := func(t *testing.T, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	}

	t.Run("sum two numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectResult(t, sum, expected)
	})

}

func ExampleAdd() {
	sum := Add(5, 6)
	fmt.Println(sum)
	// Output: 11
}
