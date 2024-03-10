package main

import "fmt"

func main() {
	nums := []int{3,2,3,4,2}
	// Output: [-2,-1,0,2,3]
	// Explanation: For index i = 0, there is 1 element in the prefix and 3 distinct elements in the suffix. Thus, diff[0] = 1 - 3 = -2.
	// For index i = 1, there are 2 distinct elements in the prefix and 3 distinct elements in the suffix. Thus, diff[1] = 2 - 3 = -1.
	// For index i = 2, there are 2 distinct elements in the prefix and 2 distinct elements in the suffix. Thus, diff[2] = 2 - 2 = 0.
	// For index i = 3, there are 3 distinct elements in the prefix and 1 distinct element in the suffix. Thus, diff[3] = 3 - 1 = 2.
	// For index i = 4, there are 3 distinct elements in the prefix and no elements in the suffix. Thus, diff[4] = 3 - 0 = 3.
	fmt.Println("Result := ",distinctDifferenceArray(nums))
}

func distinctDifferenceArray(nums []int) []int {	
	n := len(nums)
	m1,m2 := make(map[int]int),make(map[int]int)
	result := make([]int,n)

	for _,v := range nums{
		m1[v]++
	}

	for i:=0; i<n; i++ {
		m2[nums[i]]++
		m1[nums[i]]--
		if m1[nums[i]] == 0{
			delete(m1,nums[i])
		}
		result[i] = len(m2)-len(m1)
	}

	return result
}
