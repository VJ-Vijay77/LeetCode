package main

import "fmt"

func main() {
	words := []string{"leet", "code"}
	var x byte
	x = 'e'
	// Output: [0,1]

	result := findWordsContaining(words, x)
	fmt.Println("Result := ", result)
}

func findWordsContaining(words []string, x byte) []int {
	var result []int

	for i, v := range words {
		for _, t := range []byte(v) {
			if t == x {
				result = append(result, i)
				break
			}
		}
	}

	return result
}
