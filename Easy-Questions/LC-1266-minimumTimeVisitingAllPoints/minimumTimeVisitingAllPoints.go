package main

import "fmt"


func main() {
	points := [][]int{{1,1},{3,4},{-1,0}}
	// Output: 7
	// Explanation: One optimal path is [1,1] -> [2,2] -> [3,3] -> [3,4] -> [2,3] -> [1,2] -> [0,1] -> [-1,0]   
	// Time from [1,1] to [3,4] = 3 seconds 
	// Time from [3,4] to [-1,0] = 4 seconds
	// Total time = 7 seconds
	fmt.Println("Result := ",minTimeToVisitAllPoints(points))

}

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0

	prev := points[0]
	for _, p := range points[1:] {
		res += max(abs(p[0]-prev[0]), abs(p[1]-prev[1]))
		prev = p
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
