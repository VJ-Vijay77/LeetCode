package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 2
	// Output: [0,1,1]
	// Explanation:
	// 0 --> 0
	// 1 --> 1
	// 2 --> 10
	fmt.Println("Result := ", countBits(n))
}

func countBits1(n int) []int {
	var res []int
	for i := 0; i <= n; i++ {
		bin := strconv.FormatInt(int64(i), 2)
		s := strings.ReplaceAll(bin, "0", "")
		res = append(res, len(s))

	}
	return res
}

func countBits(n int) []int {
	ones := make([]int, n+1)
	for x := 0; x <= n; x++ {
        ones[x] = ones[x>>1] + x&1
	}
	return ones
}
