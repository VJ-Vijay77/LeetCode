package main

import "fmt"

func main() {
	s := []byte{'h','e','l','l','o'}
	// Output: ["o","l","l","e","h"]
	reverseString(s)
}

func reverseString(s []byte) {
left,right := 0,len(s)-1

for left < right {
	s[left],s[right] = s[right],s[left]
	left++
	right--
}
fmt.Println(string(s))

}
