package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aAbBcC"
	// Output: 2
	// Explanation:
	// From s[0] = 'a' to s[1] = 'A', there is no change of key as caps lock or shift is not counted.
	// From s[1] = 'A' to s[2] = 'b', there is a change of key.
	// From s[2] = 'b' to s[3] = 'B', there is no change of key as caps lock or shift is not counted.
	// From s[3] = 'B' to s[4] = 'c', there is a change of key.
	// From s[4] = 'c' to s[5] = 'C', there is no change of key as caps lock or shift is not counted.
	fmt.Println("Result := ",countKeyChanges(s))
}
func countKeyChanges(s string) int {
	s = strings.ToLower(s)
	count := 0
	for i:=1;i<len(s); i++{
		if s[i]!=s[i-1]{
			count++
		}
	}
	return count
}
