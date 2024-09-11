package fibonacci

import "fmt"

func Fibonacci(n int) (result []int) {
	result = make([]int, n)

	if n < 0 {
		panic("fibonnaci should have non negative input")
	}

	for i := 0; i < n; i++ {
		result[i] = fib(i)

	}

	return
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-2) + fib(n-1)
}

func main() {
	fmt.Println(Fibonacci(5))
}
