package main

import "fmt"

func main() {
	num := 14
	// Output: 6
	// Explanation:
	// Step 1) 14 is even; divide by 2 and obtain 7.
	// Step 2) 7 is odd; subtract 1 and obtain 6.
	// Step 3) 6 is even; divide by 2 and obtain 3.
	// Step 4) 3 is odd; subtract 1 and obtain 2.
	// Step 5) 2 is even; divide by 2 and obtain 1.
	// Step 6) 1 is odd; subtract 1 and obtain 0.
	fmt.Println("Result := ", numberOfSteps(num))
}

func numberOfSteps(num int) int {
	var count int
	if num %2 != 0 {
		num -= 1
		count++
	}
	for num > 0 {
		num /= 2
		if num%2 != 0 {
			num -= 1
			count++
		}
		count++
	}
	return count
}
