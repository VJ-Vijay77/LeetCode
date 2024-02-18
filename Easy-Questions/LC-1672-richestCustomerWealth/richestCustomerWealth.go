package main

import "fmt"

func main() {
	accounts := [][]int{
		{2, 8,7},
		{7,1, 3},
		{1,9, 5},
	}
	// [2,8,7],[7,1,3],[1,9,5]
	// Output: 10
	// Explanation:
	// 1st customer has wealth = 6
	// 2nd customer has wealth = 10
	// 3rd customer has wealth = 8
	// The 2nd customer is the richest with a wealth of 10.
	result := maximumWealth(accounts)
	fmt.Println("Result := ",result)
}

func maximumWealth(accounts [][]int) int {
	richestBalance := 0 
	for _,v := range accounts{
		prevBalance := 0
		for _,s := range v{
			prevBalance+=s
		}
		if prevBalance > richestBalance{
			richestBalance = prevBalance
			prevBalance = 0
		}
	}

	return richestBalance
}
