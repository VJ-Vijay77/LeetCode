package main

import "fmt"

func main() {
	nums := []int{0, 1, 4, 6, 7, 10}
	diff := 3
	// Output: 2
	// Explanation:
	// (1, 2, 4) is an arithmetic triplet because both 7 - 4 == 3 and 4 - 1 == 3.
	// (2, 4, 5) is an arithmetic triplet because both 10 - 7 == 3 and 7 - 4 == 3.
	fmt.Println(arithmeticTriplets(nums, diff))
}

func arithmeticTriplets(nums []int, diff int) int {
	count := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] == diff {
				for k := j + 1; k < len(nums); k++ {
					if nums[k] - nums[j] == diff{
						count++
					}
				}
			}
		}
	}

	return count
}
