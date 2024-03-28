package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1, 4}
	// Output: 2
	// Explanation: In this example, the minimum value is 1 and the maximum value is 4. Therefore, either 2 or 3 can be valid answers.
	fmt.Println("Result := ", findNonMinOrMax(nums))
}

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}
	sort.Ints(nums)
	return nums[1]
}
