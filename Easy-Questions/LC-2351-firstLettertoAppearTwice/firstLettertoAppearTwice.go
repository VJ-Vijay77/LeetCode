package main

import "fmt"

func main() {
	s := "abccbaacz"
	// Output: "c"
	// Explanation:
	// The letter 'a' appears on the indexes 0, 5 and 6.
	// The letter 'b' appears on the indexes 1 and 4.
	// The letter 'c' appears on the indexes 2, 3 and 7.
	// The letter 'z' appears on the index 8.
	// The letter 'c' is the first letter to appear twice, because out of all the letters the index of its second occurrence is the smallest.
	fmt.Println("Result := ",repeatedCharacter(s))
}

func repeatedCharacter(s string) string {
	rep := make(map[rune]int)
	for _,v := range s{
		rep[v]++
		if rep[v] == 2{
			return string(v)
		}
	}
	return ""
}
