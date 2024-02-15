package main

import (
	"bytes"
	"fmt"
)

func main() {
	address := "1.1.1.1"
	// Output: "1[.]1[.]1[.]1"

	result := defangIPaddr(address)
	fmt.Println("Result : ",result)
}

func defangIPaddr(address string) string {
	var result bytes.Buffer

	for _,v := range address{
		if v == '.' {
			result.WriteString("[.]")
		}else{
			result.WriteRune(v)
		}
	}
 return result.String()   
}