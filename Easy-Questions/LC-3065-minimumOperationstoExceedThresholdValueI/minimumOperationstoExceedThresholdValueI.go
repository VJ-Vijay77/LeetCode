package main

import "fmt"

func main() {
	nums := []int{2,11,10,1,3}
	k := 10
	// Output: 3
	// Explanation: After one operation, nums becomes equal to [2, 11, 10, 3].
	// After two operations, nums becomes equal to [11, 10, 3].
	// After three operations, nums becomes equal to [11, 10].
	// At this stage, all the elements of nums are greater than or equal to 10 so we can stop.
	// It can be shown that 3 is the minimum number of operations needed so that all elements of the array are greater than or equal to 10.
	fmt.Println("Result := ",minOperations(nums,k))
}

func minOperations(nums []int, k int) int {
    count := 0
	for _,v := range nums{
		if v < k {
			count++
		}
	}
	return count
}
