package interation

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectResult := func(t *testing.T, expected, repeated string) {
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	repeated := Repeat("a", 3)
	expected := "aaa"

	assertCorrectResult(t, expected, repeated)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	//Output: aaaaa
}
