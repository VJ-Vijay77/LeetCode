package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3, 2, 2, 2}
	// Output: true
	// Explanation:
	// There are 6 elements in nums, so they should be divided into 6 / 2 = 3 pairs.
	// If nums is divided into the pairs (2, 2), (3, 3), and (2, 2), it will satisfy all the conditions.
	fmt.Println("Result := ", divideArray(nums))
}

func divideArray(nums []int) bool {
	if len(nums)%2 != 0 {
		return false
	}
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	for _,v := range m{
		if v%2 != 0{
			return false
		}
	}
	return true
}
