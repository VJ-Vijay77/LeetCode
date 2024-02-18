package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"one.two.three", "four.five", "six"}
	var separator byte
	separator = '.'
	// Output: ["one","two","three","four","five","six"]
	// Explanation: In this example we split as follows:
	// "one.two.three" splits into "one", "two", "three"
	// "four.five" splits into "four", "five"
	// "six" splits into "six"
	// Hence, the resulting array is ["one","two","three","four","five","six"].

	result := splitWordsBySeparator1(words, separator)
	fmt.Println("Result := ", result)
}

func splitWordsBySeparator(words []string, separator byte) []string {

	s := string(separator)
	ansTmp, ans := []string{}, []string{}
	for _, w := range words {
		ansTmp = append(ansTmp, strings.Split(w, s)...)
	}
	for _, a := range ansTmp {
		if a != "" {
			ans = append(ans, a)
		}
	}
	return ans
}

func splitWordsBySeparator1(words []string, separator byte) []string {
	var result []string

	for _, v := range words {
		tmp := strings.Split(v, string(separator))
		for _, g := range tmp {
			if g != "" {
				result = append(result, g)
			}
		}
	}

	return result
}
