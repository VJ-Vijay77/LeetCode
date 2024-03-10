package main

import (
	"fmt"
	"strings"
)

func main() {
	patterns := []string{"a", "abc", "bc", "d"}
	word := "abc"
	// Output: 3
	// Explanation:
	// - "a" appears as a substring in "abc".
	// - "abc" appears as a substring in "abc".
	// - "bc" appears as a substring in "abc".
	// - "d" does not appear as a substring in "abc".
	// 3 of the strings in patterns appear as a substring in word.
	fmt.Println("Result := ", numOfStrings(patterns, word))
}

func numOfStrings(patterns []string, word string) int {
	count := 0
	for _, v := range patterns {
		if strings.Contains(word, v) {
			count++
		}
	}
	return count
}
