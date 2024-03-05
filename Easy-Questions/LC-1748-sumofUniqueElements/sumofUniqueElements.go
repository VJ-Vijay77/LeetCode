package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2}
	//Output: 4
	//Explanation: The unique elements are [1,3], and the sum is 4.
	fmt.Println("Result := ", sumOfUnique(nums))
}
func sumOfUnique(nums []int) int {

	m := make(map[int]int)
	var ans int

	for _, v := range nums {
		m[v]++
	}

	for value, count := range m {
		if count == 1 {
			ans += value
		}
	}

	return ans
}
