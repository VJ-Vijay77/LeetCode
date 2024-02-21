package main

import "fmt"

func main() {
	nums := []int{10,4,8,3}
	// Output: [15,1,11,22]
	// Explanation: The array leftSum is [0,10,14,22] and the array rightSum is [15,11,3,0].
	// The array answer is [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22].
	fmt.Println("Result := ",leftRightDifference(nums))
}



func leftRightDifference(nums []int) []int {
	ls := make([]int, len(nums), len(nums))
	rs := make([]int, len(nums), len(nums))
	sum := 0
	for i := 0; i < len(nums)-1; i++ {
		sum += nums[i]
		ls[i+1] = sum
	}
	sum = 0
	for i := len(nums) - 1; i > 0; i-- {
		sum += nums[i]
		rs[i-1] = sum
	}

	for i := 0; i < len(nums); i++ {
		if rs[i] > ls[i] {
			rs[i] -= ls[i]
		} else {
			rs[i] = ls[i] - rs[i]
		}
	}
	return rs
}