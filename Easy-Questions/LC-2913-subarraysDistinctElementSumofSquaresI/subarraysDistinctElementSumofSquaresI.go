package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	// Output: 15
	// Explanation: Six possible subarrays are:
	// [1]: 1 distinct value
	// [2]: 1 distinct value
	// [1]: 1 distinct value
	// [1,2]: 2 distinct values
	// [2,1]: 2 distinct values
	// [1,2,1]: 2 distinct values
	// The sum of the squares of the distinct counts in all subarrays is equal to 12 + 12 + 12 + 22 + 22 + 22 = 15.
	fmt.Println("Result := ", sumCounts(nums))
}

func sumCounts(nums []int) int {
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		s := make([]bool, 101)
		k := 0
		for j := i; j < n; j++ {
			if !s[nums[j]] {
				s[nums[j]] = true
				k++
			}
			ans += k * k
		}
	}
	return ans
}
