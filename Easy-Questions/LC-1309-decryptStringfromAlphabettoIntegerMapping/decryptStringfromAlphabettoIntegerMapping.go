package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "10#11#12"
	// Output: "jkab"
	// Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".
	fmt.Println("Result := ", freqAlphabets(s))
}

func freqAlphabets(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+2] == '#' {
            n, _ := strconv.Atoi(s[i : i+2])
			fmt.Println(n)
            res += string('a' + n - 1)
            i += 2 // 
        }else{
			n,_ := strconv.Atoi(string(s[i]))
			res += string('a'+n-1)
		}
	}
	return res
}
