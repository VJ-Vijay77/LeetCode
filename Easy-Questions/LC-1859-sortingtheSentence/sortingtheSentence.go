package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "is2 sentence4 This1 a3"
	// Output: "This is a sentence"
	// Explanation: Sort the words in s to their original positions "This1 is2 a3 sentence4", then remove the numbers.
	fmt.Println("Result := ", sortSentence(s))
}

func sortSentence(s string) string {
	var res string

	split := strings.Fields(s)

	order := make(map[string]string)

	for _, v := range split {
		order[string(v[len(v)-1])] = strings.TrimSuffix(v, string(v[len(v)-1]))
	}

	for i := 1; i <= len(s); i++ {
		res += order[strconv.Itoa(i)] + " "
	}

	return strings.TrimSpace(res)
}
