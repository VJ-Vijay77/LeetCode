package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello how are you Contestant"
	k := 4
	// Output: "Hello how are you"
	// Explanation:
	// The words in s are ["Hello", "how" "are", "you", "Contestant"].
	// The first 4 words are ["Hello", "how", "are", "you"].
	// Hence, you should return "Hello how are you".
	result := truncateSentence(s,k)
	fmt.Println("Result := ",result)
}
func truncateSentence(s string, k int) string {
	n := strings.Split(s, " ")
	return strings.Join(n[:k]," ")
}