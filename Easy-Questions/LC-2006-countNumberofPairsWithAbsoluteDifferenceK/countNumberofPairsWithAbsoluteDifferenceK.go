package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	// Output: 4
	// Explanation: The pairs with an absolute difference of 1 are:
	// - [1,2,2,1]
	// - [1,2,2,1]
	// - [1,2,2,1]
	// - [1,2,2,1]
	fmt.Println("Result := ", countKDifference(nums, k))
}

func countKDifference(nums []int, k int) int {
	noOfPairs := 0
	
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) == k {
				noOfPairs++
			}
		}
	}

	return noOfPairs
}
