package main

import "fmt"

func main() {
	nums1 := []int{4,3,2,3,1}
	 nums2 := []int{2,2,5,2,3,6}
	// Output: [3,4]
	// Explanation: We calculate the values as follows:
	// - The elements at indices 1, 2, and 3 in nums1 occur at least once in nums2. So the first value is 3.
	// - The elements at indices 0, 1, 3, and 4 in nums2 occur at least once in nums1. So the second value is 4.
	fmt.Println("Result := ",findIntersectionValues(nums1,nums2))
}


func findIntersectionValues(nums1 []int, nums2 []int) []int {
    num1 := make(map[int]bool)
    num2 := make(map[int]bool)
	
	for _,v := range nums1{
		num1[v] = true
	}
	for _,v := range nums2{
		num2[v] = true
	}
	var count1,count2 int

	for _,v := range nums1{
		if _,ok := num2[v]; ok{
			count1++
		}
	}

	for _,v := range nums2{
		if _,ok := num1[v]; ok{
			count2++
		}
	}

	return []int{count1,count2}
}