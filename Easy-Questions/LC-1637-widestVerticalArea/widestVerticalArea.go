package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{8,7},{9,9},{7,4},{9,7}}
	// Output: 1
	// Explanation: Both the red and the blue area are optimal.
	fmt.Println("Result := ",maxWidthOfVerticalArea(points))
}

func maxWidthOfVerticalArea(points [][]int) int {
	n := len(points)
	arr := make([]int, n)
	for i, point := range points {
		arr[i] = point[0]
	}
	sort.Ints(arr)
	maxWidth := arr[1] - arr[0]
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] > maxWidth {
			maxWidth = arr[i] - arr[i-1]
		}
	}
	return maxWidth
}
