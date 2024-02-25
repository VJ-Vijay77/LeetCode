package main

import "fmt"

func main() {
	s := "string"
	// Output: "rtsng"
	// Explanation:
	// After typing first character, the text on the screen is "s".
	// After the second character, the text is "st".
	// After the third character, the text is "str".
	// Since the fourth character is an 'i', the text gets reversed and becomes "rts".
	// After the fifth character, the text is "rtsn".
	// After the sixth character, the text is "rtsng".
	// Therefore, we return "rtsng".
	fmt.Println("Result := ", finalString(s))

}

func finalString(s string) string {
	var ans string

	for _,v := range s{
		if v == 'i' {
			left,right := 0,len(ans)-1
			l := []rune(ans)
			for left < right {
				l[left],l[right] = l[right],l[left]
				left++
				right--
			}
			ans = string(l)
		}else{
			ans += string(v)
		}
	}

	return ans
}
