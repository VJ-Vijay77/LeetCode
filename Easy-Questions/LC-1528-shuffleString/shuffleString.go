package main

import "fmt"

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	// Output: "leetcode"
	// Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.
	result := restoreString(s, indices)
	fmt.Println("Result := ", result)
}

func restoreString(s string, indices []int) string {
	result := make([]byte,len(s))
	
	for i,v := range indices{
		result[v] = s[i]
	}

	return string(result)
}
