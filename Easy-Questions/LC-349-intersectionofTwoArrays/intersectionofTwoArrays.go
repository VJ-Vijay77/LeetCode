package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	// Output: [9,4]
	// Explanation: [4,9] is also accepted.
	fmt.Println("Result := ", intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	elementsCount := make(map[int]bool)
	for _,v := range nums1{
		elementsCount[v]= true
	}
	fmt.Println(elementsCount)

	for _,v := range nums2{
		if elementsCount[v] {
			res = append(res, v)
			elementsCount[v] = false
		}
	}


	return res
}
