package main

import "fmt"

func main() {
	nums1 := []int{1, 1, 3, 2}
	nums2 := []int{2, 3}
	nums3 := []int{3}
	// Output: [3,2]
	// Explanation: The values that are present in at least two arrays are:
	// - 3, in all three arrays.
	// - 2, in nums1 and nums2.
	fmt.Println("Result := ",twoOutOfThree(nums1,nums2,nums3))
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
globalMap := make(map[int]int)

for _,arr := range [][]int{nums1,nums2,nums3}{
	localMap := make(map[int]struct{},len(arr))
	for _,v := range arr{
		if _,ok := localMap[v]; ok {
			continue
		}
		globalMap[v]++
		localMap[v] = struct{}{}
	}
}
var res []int
	for i,v := range globalMap{
		if v>1{
			res = append(res, i)
		}
	}

return res
}
