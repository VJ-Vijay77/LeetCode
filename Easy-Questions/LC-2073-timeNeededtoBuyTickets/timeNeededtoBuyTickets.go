package main

import "fmt"

func main() {
	tickets := []int{2, 3, 2}
	k := 2
	// Output: 6
	// Explanation:
	// - In the first pass, everyone in the line buys a ticket and the line becomes [1, 2, 1].
	// - In the second pass, everyone in the line buys a ticket and the line becomes [0, 1, 0].
	// The person at position 2 has successfully bought 2 tickets and it took 3 + 3 = 6 seconds.
	fmt.Println("Result := ",timeRequiredToBuy(tickets,k))
}

func timeRequiredToBuy(tickets []int, k int) int {
	result := 0
	c := tickets[k]
	for i, n := range tickets {
		if i == k {
			result += c
		} else if i < k {
			result += min(c, n)
		} else {
			result += min(c-1, n)
		}
	}

	return result
}
