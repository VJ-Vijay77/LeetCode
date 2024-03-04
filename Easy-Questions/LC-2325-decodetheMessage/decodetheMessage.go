package main

import (
	"fmt"
)

func main() {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	// Output: "this is a secret"
	// Explanation: The diagram above shows the substitution table.
	// It is obtained by taking the first appearance of each letter in "the quick brown fox jumps over the lazy dog".
	fmt.Println("Result := ", decodeMessage(key, message))

}
func decodeMessage(key string, message string) string {
	var result string

	letters := make(map[string]string)
	alphabet := 'a'
	for _, v := range key {
		if string(v) != " " {
			if _, ok := letters[string(v)]; !ok {
				letters[string(v)] = string(alphabet)
				alphabet++
			}
		}
	}

	for _, v := range message {
		if string(v) == " " {
			result += " "
			continue
		}
		result += letters[string(v)]
	}

	return result
}
