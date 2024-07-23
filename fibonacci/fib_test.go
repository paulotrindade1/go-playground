package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	input := 5
	want := []int{0, 1, 1, 2, 3}

	res := Fibonacci(input)

	assert.Equal(t, want, res)
}
