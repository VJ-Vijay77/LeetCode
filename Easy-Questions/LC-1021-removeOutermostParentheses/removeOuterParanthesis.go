package main

import "fmt"

func main() {
	s := "(()())(())"
	// Output: "()()()"
	// Explanation: 
	// The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
	// After removing outer parentheses of each part, this is "()()" + "()" = "()()()".
	fmt.Println("Result := ",removeOuterParentheses(s))
}
func removeOuterParentheses(s string) string {
	var result string
	count := 0
	for _,v := range s{
		if v == '(' && count > 0{
			result += string(v)
		}
		if v == '(' {
			count++
		}
		if v == ')' && count > 1{
			result += string(v)
		}
		if v == ')' {
			count--
		}
	}
	return result
}

// with switch
func removeOuterParentheses1(s string) string {
    var result string
    var index int

    for _, x := range s {
        switch x {
        case '(':
            if index > 0 {
                result += string(x)
            }
            index++
        case ')':
            index--
            if index > 0 {
                result += string(x)
            }
        }
    }

    return result
}