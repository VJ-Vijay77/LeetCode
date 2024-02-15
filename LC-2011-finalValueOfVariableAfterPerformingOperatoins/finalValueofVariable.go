package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"--X","X++","X++"}
	result := finalValueAfterOperations2(input)
	fmt.Println("Result : ",result)
}


func finalValueAfterOperations(operations []string) int {
    result := 0

	for _,v := range operations{
		if strings.HasPrefix(v, "--") || strings.HasSuffix(v, "--"){
			result--
		}else if strings.HasPrefix(v, "++") || strings.HasSuffix(v, "++"){
			result++
		}
	}
	return result 
}

// more optimised
func finalValueAfterOperations2(operations []string) int {
    X := 0
    
    // map + and - to values
    values := map[byte]int{'+': 1, '-': -1}
	fmt.Println("values map ",values)
    
    for _, operation := range operations {
        // only need to check middle character
        // and mapping of the byte
        X += values[operation[1]]
    }
    
    return X
}