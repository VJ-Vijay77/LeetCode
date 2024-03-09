package main

import "fmt"


func main(){
	nums := []int{1,2,3,4}
// Output: [2,4,4,4]
// Explanation: The first pair [1,2] means we have freq = 1 and val = 2 so we generate the array [2].
// The second pair [3,4] means we have freq = 3 and val = 4 so we generate [4,4,4].
// At the end the concatenation [2] + [4,4,4] is [2,4,4,4].
	fmt.Println("Result := ",decompressRLElist(nums))
}

func decompressRLElist(nums []int) []int {
    n := len(nums)
	var res []int
	for i:=0; i<n; i=i+2{
		freq,val := nums[i],nums[i+1]
		fmt.Println(freq, val)
		for j:=0; j<freq; j++{
			res = append(res, val)
		}
	}
	return res
}