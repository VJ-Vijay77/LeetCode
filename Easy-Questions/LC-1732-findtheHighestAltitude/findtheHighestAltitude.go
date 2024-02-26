package main

import "fmt"


func main() {
	gain := []int{-5,1,5,0,-7}
	// Output: 1
	// Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
	fmt.Println("Result := ",largestAltitude(gain))
}


func largestAltitude(gain []int) int {
	count,maxc := 0,0
    for _,v := range gain{
		count += v
		if count > maxc {
			maxc = count
		}
	}
	return maxc
}