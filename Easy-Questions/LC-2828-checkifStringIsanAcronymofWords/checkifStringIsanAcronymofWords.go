package main

import "fmt"

func main() {
	words := []string{"alice","bob","charlie"} 
	s := "abc"
	// Output: true
	// Explanation: The first character in the words "alice", "bob", and "charlie" are 'a', 'b', and 'c', respectively. Hence, s = "abc" is the acronym. 
	fmt.Println("Result := ",isAcronym(words,s))
}
func isAcronym(words []string, s string) bool {
    if len(s) != len(words){
        return false
    }
	var acr string

	for _,v := range words{
		acr += string(v[0])
	}
    return acr == s
}