package main

import "fmt"


func main() {
	names := []string{"Mary","John","Emma"} 
	heights := []int{180,165,170}
	// Output: ["Mary","Emma","John"]
	// Explanation: Mary is the tallest, followed by Emma and John.
	fmt.Println("Result := ",sortPeople(names,heights))
}


func sortPeople(names []string, heights []int) []string {

	for i:=0; i<len(heights)-1; i++{
		for j:=i+1; j<len(heights); j++{
			if heights[j] > heights[i] {
				heights[i],heights[j] = heights[j],heights[i]
				names[i],names[j] = names[j],names[i]
			} 
		}
	}
	return names
}