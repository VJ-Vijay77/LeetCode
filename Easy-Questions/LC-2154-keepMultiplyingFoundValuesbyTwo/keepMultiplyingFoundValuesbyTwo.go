package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 3, 6, 1, 12}
	original := 3
	// Output: 24
	// Explanation:
	// - 3 is found in nums. 3 is multiplied by 2 to obtain 6.
	// - 6 is found in nums. 6 is multiplied by 2 to obtain 12.
	// - 12 is found in nums. 12 is multiplied by 2 to obtain 24.
	// - 24 is not found in nums. Thus, 24 is returned.
	fmt.Println("Result := ", findFinalValue(nums, original))
}

func findFinalValue(nums []int, original int) int {
	res := 1
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == original {
			res = nums[i] * 2
			original = res
		}
	}
	if res == 1 {
		return original
	}
	return res
}
