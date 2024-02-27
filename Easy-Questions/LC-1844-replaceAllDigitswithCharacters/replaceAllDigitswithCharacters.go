package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "a1c1e1"
	// Output: "abcdef"
	// Explanation: The digits are replaced as follows:
	// - s[1] -> shift('a',1) = 'b'
	// - s[3] -> shift('c',1) = 'd'
	// - s[5] -> shift('e',1) = 'f'
	fmt.Println("Result := ",replaceDigits(s))
}
func replaceDigits(s string) string {
	str := []rune(s)
    for i,v := range s{
		if i%2!=0{
			add,_ := strconv.Atoi(string(v))
			str[i] = str[i-1]+rune(add)
		}
	}
	return string(str)
}