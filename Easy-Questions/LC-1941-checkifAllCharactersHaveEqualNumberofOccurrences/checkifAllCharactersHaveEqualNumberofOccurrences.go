package main

import "fmt"

func main() {
	s := "abacbc"
	//Output: true
	//Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.
	fmt.Println("Result := ", areOccurrencesEqual(s))
}

func areOccurrencesEqual(s string) bool {
	areEqual := make(map[string]int)
	for _, v := range s {
		areEqual[string(v)]++
	}
	check := areEqual[string(s[0])]
	for _, v := range areEqual {
		if v != check {
			return false
		}
	}
	return true
}
