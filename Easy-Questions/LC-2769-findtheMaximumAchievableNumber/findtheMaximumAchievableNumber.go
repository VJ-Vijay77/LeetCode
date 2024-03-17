package main

import "fmt"

func main() {
	num := 4
	t := 1
	// Output: 6
	// Explanation: The maximum achievable number is x = 6; it can become equal to num after performing this operation:
	// 1- Decrease x by 1, and increase num by 1. Now, x = 5 and num = 5. 
	// It can be proven that there is no achievable number larger than 6.
	fmt.Println("Result := ",theMaximumAchievableX(num,t))
}

func theMaximumAchievableX(num int, t int) int {
    // num+t = x-t
	// num + t + t = x
	return num + t + t
}