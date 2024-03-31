package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4,-1,0,3,10}
	// Output: [0,1,9,16,100]
	// Explanation: After squaring, the array becomes [16,1,0,9,100].
	// After sorting, it becomes [0,1,9,16,100].
	fmt.Println("Result := ",sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	arr := make([]int,len(nums))
	for i,v := range nums{
		arr[i] = v*v
	}
	sort.Ints(arr)
	return arr
}
