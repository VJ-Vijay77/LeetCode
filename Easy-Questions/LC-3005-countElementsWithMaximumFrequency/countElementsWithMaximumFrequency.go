package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 1, 4}
	// Output: 4
	// Explanation: The elements 1 and 2 have a frequency of 2 which is the maximum frequency in the array.
	// So the number of elements in the array with maximum frequency is 4.
	fmt.Println("Result := ", maxFrequencyElements(nums))
}
func maxFrequencyElements(nums []int) int {
	count := 0
	freq := make(map[int]int)
	maxCount := 0
	for _, v := range nums {
		freq[v]++
		if freq[v] > maxCount {
			maxCount = freq[v]
		}
	}
	for _, v := range freq {
		if v == maxCount{
			count+=maxCount
		}
	}
	return count
}
