package main

import "fmt"

func main() {
	nums := []int{3,1,2,4}
	// Output: [2,4,3,1]
	// Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
	fmt.Println("Result := ",sortArrayByParity(nums))
}

func sortArrayByParity(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	var res []int
	var odd []int
	for _,v := range nums{
		if v % 2 == 0 {
			res = append(res, v)
		}else{
			odd = append(odd, v)
		}
	}
	res = append(res, odd...)
	return res
}
