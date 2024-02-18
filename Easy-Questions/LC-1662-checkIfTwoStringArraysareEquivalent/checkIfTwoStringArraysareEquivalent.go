package main

import "fmt"

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	// Output: true
	// Explanation:
	// word1 represents string "ab" + "c" -> "abc"
	// word2 represents string "a" + "bc" -> "abc"
	// The strings are the same, so return true.
	result := arrayStringsAreEqual(word1, word2)
	fmt.Println("Result := ", result)
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var isEqual bool
	var wordfull1, wordfull2 string

	for _, v := range word1 {
		wordfull1 += v
	}
	for _, v := range word2 {
		wordfull2 += v
	}
	if wordfull1 == wordfull2 {
		isEqual = true
	}
	return isEqual
}
