package main

import "fmt"

func main() {
	command := "G()(al)"
	// Output: "Goal"
	// Explanation: The Goal Parser interprets the command as follows:
	// G -> G
	// () -> o
	// (al) -> al
	// The final concatenated result is "Goal".
	resutl := interpret(command)
	fmt.Println("Result := ",resutl)
}

func interpret(command string) string {
    var result string
	for i:=0; i<len(command); i++ {
		switch (command[i]) {
		case 'G':
			result+= "G"
		case '(':
			if command[i] == '(' && command[i+1] == ')' {
				result+= "o"
			} else if command[i] == '(' && command[i+1] == 'a' {
				result+= "al"
			}	
		}
	}

	return result
}