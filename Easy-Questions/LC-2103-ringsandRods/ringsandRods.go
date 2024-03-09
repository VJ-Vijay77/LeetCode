package main

import (
	"fmt"
	"strconv"
)

func main() {
	rings := "B0B6G0R6R0R6G9"
	// Output: 1
	// Explanation:
	// - The rod labeled 0 holds 3 rings with all colors: red, green, and blue.
	// - The rod labeled 6 holds 3 rings, but it only has red and blue.
	// - The rod labeled 9 holds only a green ring.
	// Thus, the number of rods with all three colors is 1.
	fmt.Println("Result := ", countPoints(rings))
}

func countPoints(rings string) int {
	r := make([]bool, 10)
	g := make([]bool, 10)
	b := make([]bool, 10)
	count := 0
	for i := 0; i < len(rings)-1; i = i + 2 {
		if rings[i] == 'R' {
			a, _ := strconv.Atoi(string(rings[i+1]))
			r[a] = true
		} else if rings[i] == 'G' {
			b, _ := strconv.Atoi(string(rings[i+1]))
			g[b] = true
		} else if rings[i] == 'B' {
			c, _ := strconv.Atoi(string(rings[i+1]))
			b[c] = true
		}
	}
	for i := 0; i < 10; i++ {
		if r[i] == true && g[i] == true && b[i] == true {
			count++
		}
	}
	return count
}
