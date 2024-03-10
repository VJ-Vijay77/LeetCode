package main

import "fmt"

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// Output: 25
	// Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
	// Notice that element mat[1][1] = 5 is counted only once.
	fmt.Println("Result := ",diagonalSum(mat))
}

func diagonalSum(mat [][]int) int {
	sum,n := 0,len(mat)

	for i:=0; i<n; i++{
		sum += mat[i][i]
		if i != n-i-1{
			sum += mat[i][n-1-i]
		}
	}

	return sum
}
