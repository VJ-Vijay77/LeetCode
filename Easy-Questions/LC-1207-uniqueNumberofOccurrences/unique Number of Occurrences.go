package main

import "fmt"

func main() {
	arr := []int{-5, -6, 2, 6, -5, -7, 5}
	//Output: true
	//Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
	fmt.Println("Result := ", uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	countOfInt := make(map[int]int)
	seen := make(map[int]struct{})
	for _, v := range arr {
		countOfInt[v]++
	}

	for _, v := range countOfInt {
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = struct{}{}
	}
	return true
}
