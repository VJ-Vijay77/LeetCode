package main

import "fmt"




func main() {
	arr := []int{1,4,2,5,3}
	// Output: 58
	// Explanation: The odd-length subarrays of arr and their sums are:
	// [1] = 1
	// [4] = 4
	// [2] = 2
	// [5] = 5
	// [3] = 3
	// [1,4,2] = 7
	// [4,2,5] = 11
	// [2,5,3] = 10
	// [1,4,2,5,3] = 15
	// If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
	fmt.Println("Result := ",sumOddLengthSubarrays(arr))

}

func sumOddLengthSubarrays(arr []int) int {
	res := 0
	for index, value := range arr {
		cnt := ((index+1)*(len(arr)-index) + 1) / 2
		res += value * cnt
	}
	return res
}