package main

import "fmt"

func main() {
	n := 7
	// Output: 21
	// Explanation: Numbers in the range [1, 7] that are divisible by 3, 5, or 7 are 3, 5, 6, 7. The sum of these numbers is 21.
	fmt.Println("Result := ", sumOfMultiples(n))
}

func sumOfMultiples(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			res += i
		}
	}
	return res
}
