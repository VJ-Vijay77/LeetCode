package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 234
	// Output: 15
	// Explanation:
	// Product of digits = 2 * 3 * 4 = 24
	// Sum of digits = 2 + 3 + 4 = 9
	// Result = 24 - 9 = 15
	fmt.Println("Result := ", subtractProductAndSum(n))
}

func subtractProductAndSum(n int) int {
	mult,add := 1,0

	s := strconv.Itoa(n)
	for _,v := range s{
		dig,_ := strconv.Atoi(string(v))
		mult *= dig
		add += dig
	}
	return mult - add
}
