package main

import (
	"fmt"
	"sort"
)

func main() {
	heights := []int{1,1,4,2,1,3}
	// Output: 3
	// Explanation: 
	// heights:  [1,1,4,2,1,3]
	// expected: [1,1,1,2,3,4]
	// Indices 2, 4, and 5 do not match.
	fmt.Println("Result := ",heightChecker(heights))
}

func heightChecker(heights []int) int {
	var dup []int
	for _,v := range heights{
		dup = append(dup, v)
	}
	sort.Ints(heights)
	count := 0
	for i:=0; i<len(heights); i++ {
		if heights[i] != dup[i]{
			count++
		}

	}
	return count
}
