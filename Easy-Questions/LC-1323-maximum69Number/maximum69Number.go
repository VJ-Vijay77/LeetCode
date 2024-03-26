package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := 9669
	// Output: 9969
	// Explanation:
	// Changing the first digit results in 6669.
	// Changing the second digit results in 9969.
	// Changing the third digit results in 9699.
	// Changing the fourth digit results in 9666.
	// The maximum number is 9969.
	fmt.Println("Result := ", maximum69Number(num))
}

func maximum69Number(num int) int {
	r, _ := strconv.Atoi((strings.Replace(strconv.Itoa(num), "6", "9", 1)))
	return r
}
