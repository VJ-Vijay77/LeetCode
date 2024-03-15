package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaaabbbbcccc"
	// Output: "abccbaabccba"
	// Explanation: After steps 1, 2 and 3 of the first iteration, result = "abc"
	// After steps 4, 5 and 6 of the first iteration, result = "abccba"
	// First iteration is done. Now s = "aabbcc" and we go back to step 1
	// After steps 1, 2 and 3 of the second iteration, result = "abccbaabc"
	// After steps 4, 5 and 6 of the second iteration, result = "abccbaabccba"
	// Example 2:
	fmt.Println("Result := ",sortString(s))
}

func sortString(s string) string {
	char := make(map[rune]int)

	for _,v:= range s{
		char[v]++
	}

	count := len(s)
	var res strings.Builder

	for count!=0{
		for i:='a'; i<='z'; i++{
			if char[i]!=0{
				char[i]--
				count--
				res.WriteRune(i)
			}
		}
		for i:='z'; i>='a'; i--{
			if char[i]!=0{
				char[i]--
				count--
				res.WriteRune(i)
			}
		}
	}
return res.String()
}
