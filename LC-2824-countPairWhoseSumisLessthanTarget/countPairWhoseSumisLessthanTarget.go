package main

import "fmt"

func main() {
	nums := []int{-1, 1, 2, 3, 1}
	target := 2
	// Output: 3
	// Explanation: There are 3 pairs of indices that satisfy the conditions in the statement:
	// - (0, 1) since 0 < 1 and nums[0] + nums[1] = 0 < target
	// - (0, 2) since 0 < 2 and nums[0] + nums[2] = 1 < target
	// - (0, 4) since 0 < 4 and nums[0] + nums[4] = 0 < target
	// Note that (0, 3) is not counted since nums[0] + nums[3] is not strictly less than the target.
	result := countPairs(nums, target)
	fmt.Println("Result  := ",result)
}

func countPairs(nums []int, target int) int {
	var result int

	for i:=0; i<len(nums); i++ {
		for j:=i+1; j<len(nums); j++ {
			if nums[i]+nums[j] < target{
				result++
			}
		}
	}

	return result
}
