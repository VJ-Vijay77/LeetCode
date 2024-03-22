package main

import "fmt"


func main(){
	n := 7
	// Output: 6
	// Explanation: Details of the tournament: 
	// - 1st Round: Teams = 7, Matches = 3, and 4 teams advance.
	// - 2nd Round: Teams = 4, Matches = 2, and 2 teams advance.
	// - 3rd Round: Teams = 2, Matches = 1, and 1 team is declared the winner.
	// Total number of matches = 3 + 2 + 1 = 6.
	fmt.Println("Result := ",numberOfMatches(n))
}

func numberOfMatches(n int) int {
    matches := 0
	count := 0
	for n>0{
		if n%2 == 0 {
			matches = n/2
		}else{
			matches = (n-1)/2
		}
		count += matches
		n -= n/2
		if n == 1 || n == 0{
			break
		}
	} 
	return count
}