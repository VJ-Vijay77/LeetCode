package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{2, 4}
	// Output: 2
	// Explanation: There are two common elements in the array 2 and 3 out of which 2 is the smallest, so 2 is returned.
	fmt.Println("Result := ", getCommon(nums1, nums2))
}

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		fmt.Println("i=",i," j=",j)
		fmt.Println("nums1[i]=",nums1[i], " nums2[j]=",nums2[j])
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			fmt.Println("i++")
			i++
		} else {
			fmt.Println("j++")
			j++
		}
	}
	return -1
}
