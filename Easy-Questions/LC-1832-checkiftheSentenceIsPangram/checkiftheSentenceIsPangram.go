package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	// Output: true
	// Explanation: sentence contains at least one of every letter of the English alphabet.
	fmt.Println("Result := ", checkIfPangram(sentence))
}

func checkIfPangram(sentence string) bool {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	res := make(map[string]string)
	check := 0
	for _, v := range []rune(sentence) {
		for _, k := range alphabet {
			if v == k {
				if _, ok := res[string(k)]; !ok {
					res[string(k)] = string(v)
					check++
				}
			}
		}
	}
	return check == 26
}

//well optimised
func checkIfPangram2(sentence string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, char := range alphabet {
		if !strings.Contains(sentence, string(char)) {
			return false
		}
	}

	return true
}
