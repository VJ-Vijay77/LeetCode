package main

import "fmt"





func main() {
	var nums = []int{0,2,1,5,3,4}
	result := buildArray(nums)
	fmt.Println("Result : ",result)
}

func buildArray(nums []int) []int{
	var ans = make([]int,len(nums))

	for i := range nums{
		ans[i] = nums[nums[i]]
	}
	
	return ans
}