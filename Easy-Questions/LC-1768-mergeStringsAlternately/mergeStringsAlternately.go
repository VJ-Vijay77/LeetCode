package main

import "fmt"

func main() {
	word1 := "ab"
	word2 := "pqrs"
	// Output: "apbqcr"
	// Explanation: The merged string will be merged as so:
	// word1:  a   b   c
	// word2:    p   q   r
	// merged: a p b q c r
	fmt.Println("Resulut := ", mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {
	var newstring string
	count := 1

	for i := 0; i < len(word1); i++ {

		newstring += string(word1[i])
		newstring += string(word2[i])

		if count == len(word1) {
			newstring += string(word2[i+1:])
			break
		}

		if count == len(word2) {
			newstring += string(word1[i+1:])
			break
		}

		count++
	}

	return newstring
}
