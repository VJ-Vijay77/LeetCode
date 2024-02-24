package main

import (
	"fmt"
	"strings"
)

func main() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	// Output: 2
	// Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.
	fmt.Println("Result := ", countConsistentStrings(allowed, words))
}
func countConsistentStrings(allowed string, words []string) int {
	count := 0

	for _, v := range words {
		for i := 0; i < len(allowed); i++ {
			v = strings.ReplaceAll(v, string(allowed[i]), "")
		}
		if len(v) == 0 {
			count++
		}
	}
	return count
}
