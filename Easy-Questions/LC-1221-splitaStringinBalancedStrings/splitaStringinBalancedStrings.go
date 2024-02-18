package main

import (
	"fmt"
)

func main() {
	s := "RLRRLLRLRL"
	// Output: 4
	// Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.
	result := balancedStringSplit(s)
	fmt.Println("Result := ", result)
}

func balancedStringSplit(s string) int {
	numOfBalancedString := 0
	Lcount := 0
	Rcount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			Lcount++
		} else {
			Rcount++
		}
		if Lcount == Rcount {
			numOfBalancedString++
		}
	}
	return numOfBalancedString
}

// little more optimised solution
func balancedStringSplit2(s string) int {
    var result, counter int
    for _, char := range s {
        if char == 'R' {
            counter++
        } else {
            counter--
        }
        if counter == 0 {
            result++
        }
    }
    return result
}