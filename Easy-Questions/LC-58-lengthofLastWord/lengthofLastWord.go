package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "   fly me   to   the moon  "
	// Output: 4
	// Explanation: The last word is "moon" with length 4.
	fmt.Println("Resul := ", lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	f := strings.Split(strings.TrimSpace(s)," ")
	return len(f[len(f)-1])
}
