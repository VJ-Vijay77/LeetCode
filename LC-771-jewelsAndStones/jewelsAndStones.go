package main

import "fmt"

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	//Output: 3
	result := numJewelsInStones(jewels, stones)
	fmt.Println("Result := ", result)
}

func numJewelsInStones(jewels string, stones string) int {
	result := 0

	jewelCollection := make(map[rune]struct{})

	for _, v := range jewels {
		jewelCollection[v] = struct{}{}
	}

	for _, v := range stones {
		if _, ok := jewelCollection[v]; ok {
			result++
		}
	}

	return result
}
