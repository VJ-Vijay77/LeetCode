package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3,4,5,2}
	// Output: 12 
	// Explanation: If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value, that is, (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12. 
	fmt.Println("Result := ",maxProduct(nums))
}

func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1]-1)*(nums[len(nums)-2]-1)
}
