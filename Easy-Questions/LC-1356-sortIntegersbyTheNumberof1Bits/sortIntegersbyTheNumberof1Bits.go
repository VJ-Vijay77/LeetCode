package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// Output: [0,1,2,4,8,3,5,6,7]
	// Explantion: [0] is the only integer with 0 bits.
	// [1,2,4,8] all have 1 bit.
	// [3,5,6] have 2 bits.
	// [7] has 3 bits.
	// The sorted array by bits is [0,1,2,4,8,3,5,6,7]
	fmt.Println("Result := ", sortByBits(arr))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		countI := bits.OnesCount(uint(arr[i]))
		countJ := bits.OnesCount(uint(arr[j]))
		return countI < countJ || (countI == countJ && arr[i] < arr[j])
	})
	return arr
}
