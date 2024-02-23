package main

import (
	"fmt"
	"strconv"
)

func main() {

	nums := []int{1,15,6,3}
// Output: 9
// Explanation: 
// The element sum of nums is 1 + 15 + 6 + 3 = 25.
// The digit sum of nums is 1 + 1 + 5 + 6 + 3 = 16.
// The absolute difference between the element sum and digit sum is |25 - 16| = 9.
	fmt.Println("Result := ",differenceOfSum(nums))
}

func differenceOfSum(nums []int) int {
	var diff,elementSum,DigitSum int
	
	for _,v := range nums{
		elementSum += v
	}
	for _,v := range nums{
		add := 0
		str := strconv.Itoa(v)
		if len(str) > 1 {
			for _,l := range []rune(str){
				j,_ := strconv.Atoi(string(l))
				add += j
			}
		}else{
			add = v
		}
		DigitSum += add
	}
	if elementSum > DigitSum{
		diff = elementSum - DigitSum
	}else {
		diff = DigitSum - elementSum
	}
	return diff
}
