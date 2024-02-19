package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	s := "K1:L2"
	// Output: ["K1","K2","L1","L2"]
	// Explanation:
	// The above diagram shows the cells which should be present in the list.
	// The red arrows denote the order in which the cells should be presented.
	fmt.Println("Result := ",cellsInRange(s))
}

func cellsInRange1(s string) []string {
	a := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var r []string
	firstLetter := string(s[0])
	secondLetter := string(s[3])
	firstIndex, _ := strconv.Atoi(string(s[1]))
	secondIndex, _ := strconv.Atoi(string(s[4]))

	indexFLetter := slices.Index(a, firstLetter)
	indexSLetter := slices.Index(a, secondLetter)
	fmt.Println("indexF Letter",indexFLetter)
	fmt.Println("indexS Letter",indexSLetter)

	h := a[indexFLetter : indexSLetter+1]
	fmt.Println("h := ",h)
	if firstIndex == secondIndex {
		for i := range h {
			add := strconv.Itoa(firstIndex)
			r = append(r, h[i]+add)
		}
	} else {
		for _, letter := range h {
			for i := firstIndex; i <= secondIndex; i++ {
				add := strconv.Itoa(i)
				r = append(r, letter+add)
			}
		}
	}

	return r
}


func cellsInRange(s string) []string {
    result := []string{}
    //column parse
    rowStart := s[1]
    rowEnd := s[4]
    colStart := s[0]
    colEnd := s[3]

    for colStart <= colEnd {
        for rowStart <= rowEnd {
            newStr := []byte{colStart,rowStart}
			fmt.Println(string(newStr))
            result = append(result, string(newStr))
            rowStart++
        }
        colStart++
        rowStart = s[1]
    }
    return result
}