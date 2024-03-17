package main

import (
	"fmt"
	"sort"
)

func main() {
	num := 2932
	// Output: 52
	// Explanation: Some possible pairs [new1, new2] are [29, 23], [223, 9], etc.
	// The minimum sum can be obtained by the pair [29, 23]: 29 + 23 = 52.
	fmt.Println("Result := ", minimumSum(num))
}

func minimumSum(num int) int {
	digit := []int{}

	for num > 0 {
		digit = append(digit, num%10)
		num /= 10
	}

	sort.Ints(digit)
	new1 := digit[0]*10 + digit[2]
	new2 := digit[1] * 10 + digit[3]

	return new1 + new2
}
