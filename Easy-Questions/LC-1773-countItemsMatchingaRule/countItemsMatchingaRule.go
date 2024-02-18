package main

import (
	"fmt"
)

func main() {
	// items := [][]string{
	// 	{"phone", "blue", "pixel"},
	// 	{"computer", "silver", "lenovo"},
	// 	{"phone", "gold", "iphone"},
	// }
	// ruleKey := "color"
	// ruleValue := "silver"
	// type | color | name
	// Output: 1
	// Explanation: There is only one item matching the given rule, which is ["computer","silver","lenovo"].

	items := [][]string{
		{"phone", "blue", "pixel"},
		{"computer", "silver", "phone"},
		{"phone", "gold", "iphone"},
	}
	ruleKey := "type"
	ruleValue := "phone"
	// Output: 2
	// Explanation: There are only two items matching the given rule, which are ["phone","blue","pixel"] and ["phone","gold","iphone"]. Note that the item ["computer","silver","phone"] does not match.
	result := countMatches(items, ruleKey, ruleValue)
	fmt.Println("Result := ", result)
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var result int
	var order = map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}
	for _, v := range items {
		for j, k := range v {
			if k == ruleValue && order[ruleKey] == j {
				result++
			}
		}
	}
	return result
}
