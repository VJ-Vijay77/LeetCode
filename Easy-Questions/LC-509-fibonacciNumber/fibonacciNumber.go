package main

import (
	"fmt"
)

func main() {
	n := 2
	// Output: 1
	// Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
	fmt.Println("Result := ", fib(n))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
