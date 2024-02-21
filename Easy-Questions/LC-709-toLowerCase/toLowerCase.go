package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello"
	// Output: "hello"

	fmt.Println("Result := ",s)
}

func toLowerCase(s string) string {
return strings.ToLower(s)
}
