package main

import "fmt"


func main(){
	nums := []int{1,2,3,3}
// Output: 3
	fmt.Println("Result := ",repeatedNTimes(nums))
}

func repeatedNTimes(nums []int) int {
    n := len(nums)/2
	countMap := make(map[int]int)

	for _,v := range nums{
		countMap[v]++
		if countMap[v] == n {
			return v		
		}
	}
	return 0
}