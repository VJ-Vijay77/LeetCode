package main

import (
	"fmt"
	"sort"
)


func main(){
	nums := []int{1,4,3,2}
// Output: 4
// Explanation: All possible pairings (ignoring the ordering of elements) are:
// 1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
// 2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
// 3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
// So the maximum possible sum is 4.
fmt.Println("Ressul := ",arrayPairSum(nums))
}


func arrayPairSum(nums []int) int {
    sort.Ints(nums)
	var res int
	for i:=0; i<len(nums); i+=2{
		res += nums[i]
	}
	return res
}