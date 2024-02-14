package main

import "fmt"


// asked to reorder array elements based on their index value
// array of 0,2,1
// result 0,1,2    because 0 comes under 0th index, arrays index 1 element is 2, so second element needed to be in the next element in result array, like wise 
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