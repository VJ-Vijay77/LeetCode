package main

import "fmt"

func main() {

	n := 10
	m := 3
	// Output: 19
	// Explanation: In the given example:
	// - Integers in the range [1, 10] that are not divisible by 3 are [1,2,4,5,7,8,10], num1 is the sum of those integers = 37.
	// - Integers in the range [1, 10] that are divisible by 3 are [3,6,9], num2 is the sum of those integers = 18.
	// We return 37 - 18 = 19 as the answer.
	fmt.Println("Result := ",differenceOfSums(n,m))
}

func differenceOfSums(n int, m int) int {
divisible := make(map[int]int)
notDivisible := make(map[int]int)

for i:=1; i<=n; i++ {
	if i % m == 0 {
		divisible[1] += i
	}else{
		notDivisible[1] += i
	}
}

return  notDivisible[1] - divisible[1]
}
