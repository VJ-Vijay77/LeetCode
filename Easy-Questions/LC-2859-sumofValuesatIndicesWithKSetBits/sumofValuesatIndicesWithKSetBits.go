package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{5, 10, 1, 5, 2}
	k := 1
	// Output: 13
	// Explanation: The binary representation of the indices are:
	// 0 = 000
	// 1 = 001
	// 2 = 010
	// 3 = 011
	// 4 = 100
	// Indices 1, 2, and 4 have k = 1 set bits in their binary representation.
	// Hence, the answer is nums[1] + nums[2] + nums[4] = 13.
	s := strconv.FormatInt(int64(2), 2)
	l := strings.Count(s, "1")
	fmt.Println(l)
	result := sumIndicesWithKSetBits(nums, k)
	fmt.Println("Result := ", result)

}

func sumIndicesWithKSetBits(nums []int, k int) int {
	ans := 0
	for i, v := range nums {
		if strings.Count(strconv.FormatInt(int64(i), 2), "1") == k {
			ans += v
		}
	}
	return ans
}
