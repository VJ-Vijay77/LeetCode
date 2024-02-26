package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 6, 2, 7, 4}
	// Output: 34
	// Explanation: We can choose indices 1 and 3 for the first pair (6, 7) and indices 2 and 4 for the second pair (2, 4).
	// The product difference is (6 * 7) - (2 * 4) = 34.
	fmt.Println(maxProductDifference(nums))
}

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	firstPair1 := nums[len(nums)-1]
	firstPair2 := nums[len(nums)-2]
	lastPair1 := nums[0]
	lastPair2 := nums[1]
	
	return (firstPair1 * firstPair2) - (lastPair1 * lastPair2)
}
