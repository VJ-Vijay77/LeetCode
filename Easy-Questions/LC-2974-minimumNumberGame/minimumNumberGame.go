package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 4, 2, 3}
	// Output: [3,2,5,4]
	// Explanation: In round one, first Alice removes 2 and then Bob removes 3. Then in arr firstly Bob appends 3 and then Alice appends 2. So arr = [3,2].
	// At the begining of round two, nums = [5,4]. Now, first Alice removes 4 and then Bob removes 5. Then both append in arr which becomes [3,2,5,4].
	fmt.Println("Result := ", numberGame(nums))
}
func numberGame(nums []int) []int {
	sort.Ints(nums)
	temp := make([]int, 2)
	var result []int
	for i := 0; i <= len(nums)-2; i = i + 2 {
		temp[1] = nums[i]
		temp[0] = nums[i+1]
		result = append(result, temp...)
	}
	return result
}

// more optimised
func numberGame1(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}
