package main

import "fmt"

func main() {
	words := []string{"cd","ac","dc","ca","zz"}
	// Output: 2
	// Explanation: In this example, we can form 2 pair of strings in the following way:
	// - We pair the 0th string with the 2nd string, as the reversed string of word[0] is "dc" and is equal to words[2].
	// - We pair the 1st string with the 3rd string, as the reversed string of word[1] is "ca" and is equal to words[3].
	// It can be proven that 2 is the maximum number of pairs that can be formed.
	fmt.Println("Result := ",maximumNumberOfStringPairs(words))
}
func maximumNumberOfStringPairs(words []string) int {
	count:=0
	for i:=0; i<len(words)-1; i++{
		for j:=i+1; j<len(words); j++{
			if words[i][0]==words[j][1] && words[i][1]==words[j][0]{
				count++
			} 
		}
	}
	return count
}
