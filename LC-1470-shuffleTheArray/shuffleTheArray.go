package main

import "fmt"

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	//Output: [2,3,5,4,1,7]
	result := shuffle(nums, n)
	fmt.Println("Result := ", result)
}

func shuffle(nums []int, n int) []int {
	var ans = make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i], nums[i+n])
	}
	return ans
}