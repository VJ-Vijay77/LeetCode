package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
	fmt.Println("Result := ",reverseWords(s))
}


func reverseWords(s string) string {
	words := strings.Fields(s)
	var res string
	for _,v := range words{
		k := []rune(v)
		for i,j:=0,len(v)-1; i<j; i,j=i+1,j-1{
			k[i],k[j] = k[j],k[i]
		}
		res += string(k)+" "
	}

	return strings.TrimSpace(res)
}