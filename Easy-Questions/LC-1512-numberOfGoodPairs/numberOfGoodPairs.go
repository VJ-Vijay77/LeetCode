package main

import "fmt"


func main() {
	input := []int{1,2,3,1,1,3}
	//expected result = 4  
	//[(0,3),(0,4),(3,4),(2,5)]
	result := numIdenticalPairs(input)
	fmt.Println("Result : ",result)
}

func numIdenticalPairs(nums []int) int {
	result := 0

	for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++{
			if nums[i] == nums[j]{
				result++
			}
		}
	}
	return result
}