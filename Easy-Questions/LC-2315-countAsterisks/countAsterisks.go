package main

import "fmt"

func main() {
	s := "l|*e*et|c**o|*de|"
	// Output: 2
	// Explanation: The considered characters are underlined: "l|*e*et|c**o|*de|".
	// The characters between the first and second '|' are excluded from the answer.
	// Also, the characters between the third and fourth '|' are excluded from the answer.
	// There are 2 asterisks considered. Therefore, we return 2.
	fmt.Println("Result := ",countAsterisks(s))

}
func countAsterisks(s string) int {
    bar,count := 0,0
	for _,v := range s{
		if v == '|' {
			bar++
		}
		if v == '*' && bar % 2 == 0{
			count++
		}
	}
	return count
}