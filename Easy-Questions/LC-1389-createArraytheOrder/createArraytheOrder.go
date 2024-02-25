package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	// Output: [0,4,1,3,2]
	// Explanation:
	// nums       index     target
	// 0            0        [0]
	// 1            1        [0,1]
	// 2            2        [0,1,2]
	// 3            2        [0,1,3,2]
	// 4            1        [0,4,1,3,2]

	fmt.Println("Result := ", createTargetArray(nums, index))
}
func createTargetArray(nums []int, index []int) []int {
	var target []int
	for i, num := range nums {
		idx := index[i]
		target = append(target[:idx], append([]int{num}, target[idx:]...)...)
	}
	return target
}
