package main

import (
	"fmt"
	"strings"
)


func main(){
	s := "book"
	// Output: true
	// Explanation: a
	fmt.Println("Result := ",halvesAreAlike(s))
}

func halvesAreAlike(s string) bool {
    n := (len(s)/2)
	vowel := map[string]int{
		"a":0,
		"e":0,
		"i":0,
		"o":0,
		"u":0,
	}
	first := strings.ToLower( string(s[:n]))
	second := strings.ToLower( string(s[n:]))
	firstC,secondC := 0,0

	for i:=0; i<n; i++{
		if _,ok := vowel[string(first[i])]; ok{
			firstC++
		}
		if _,ok := vowel[string(second[i])]; ok{
			secondC++
		}
	}

	return firstC == secondC
}