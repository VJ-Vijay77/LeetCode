package main

import "fmt"

func main() {
	nums := []int{8,1,2,2,3}
	// Output: [4,0,1,1,3]
	// Explanation: 
	// For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3). 
	// For nums[1]=1 does not exist any smaller number than it.
	// For nums[2]=2 there exist one smaller number than it (1). 
	// For nums[3]=2 there exist one smaller number than it (1). 
	// For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).
	result := smallerNumbersThanCurrent(nums)
	fmt.Println("Result := ",result)
}

func smallerNumbersThanCurrent(nums []int) []int {
	var ans []int
	for i:=0; i<len(nums); i++ {
		counter := 0
		for j:=0; j<len(nums); j++{
			if i!=j && nums[i] > nums[j] {
				counter++
			}
		}
		ans = append(ans, counter)
	}
	return ans
}
