package main

import "fmt"

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	// Output: 3
	// Explanation: Digit 8 is inside of 3 nested parentheses in the string.
	fmt.Println("Result := ", maxDepth(s))
}

func maxDepth(s string) int {
	maxDepth, count := 0, 0
	for _, v := range s {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
		}
		if count > maxDepth{
			maxDepth = count
		}
	}
	return maxDepth
}
