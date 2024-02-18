package main

import (
	"fmt"
	"strings"
)


func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	// Output: 6
	// Explanation: 
	// - The first sentence, "alice and bob love leetcode", has 5 words in total.
	// - The second sentence, "i think so too", has 4 words in total.
	// - The third sentence, "this is great thanks very much", has 6 words in total.
	// Thus, the maximum number of words in a single sentence comes from the third sentence, which has 6 words.
	result := mostWordsFound(sentences)
	fmt.Println("Result := ",result)
}

func mostWordsFound(sentences []string) int {
	maxCount := 0
	
	for _,v := range sentences{
		split := strings.Split(v, " ")
		if len(split) > maxCount {
			maxCount = len(split)
		}

	}

	return maxCount
}