package main

import "fmt"

func main() {
	word := "abcdefd"
	ch := byte('d')
	// Output: "dcbaefd"
	// Explanation: The first occurrence of "d" is at index 3. 
	// Reverse the part of word from 0 to 3 (inclusive), the resulting string is "dcbaefd".
	fmt.Println("Result := ",reversePrefix(word,ch))
}


func reversePrefix(word string, ch byte) string {
	a := []byte(word)

    for i := 0; i < len(word); i++ {
        if word[i] == ch {
            for j := 0; j <= i/2; j++ {
                temp := a[j]
                a[j] = a[i-j]
                a[i-j] = temp
            }

            return string(a)
        }
    }

    return word
}