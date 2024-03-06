package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	// Output: [[1,3],[4,6]]
	// Explanation:
	// For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
	// For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].
	fmt.Println("Result := ", findDifference(nums1, nums2))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var res [][]int
    map1 := make(map[int]struct{})
    map2 := make(map[int]struct{})
    
    for _, v := range nums1 {
        map1[v] = struct{}{}
    }
    
    for _, v := range nums2 {
        map2[v] = struct{}{}
    }

    res = append(res, []int{}) // Initialize the first slice
    res = append(res, []int{}) // Initialize the second slice

    addedValues := make(map[int]struct{}) // Map to keep track of added values

    for _, v := range nums1 {
        if _, ok := map2[v]; !ok {
            if _, added := addedValues[v]; !added {
                res[0] = append(res[0], v)
                addedValues[v] = struct{}{}
            }
        }
    }

    for _, v := range nums2 {
        if _, ok := map1[v]; !ok {
            if _, added := addedValues[v]; !added {
                res[1] = append(res[1], v)
                addedValues[v] = struct{}{}
            }
        }
    }

    return res
}
