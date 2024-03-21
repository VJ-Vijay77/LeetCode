package main

import "fmt"

func main() {
	n := 5
	// Output: 10
	// Explanation: The smallest multiple of both 5 and 2 is 10.
	fmt.Println("Ressult := ",smallestEvenMultiple(n))
}

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}
