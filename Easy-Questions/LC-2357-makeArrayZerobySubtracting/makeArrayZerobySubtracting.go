package main

import "fmt"

func main() {
	nums := []int{1,5,0,3,5}
// Output: 3
// Explanation:
// In the first operation, choose x = 1. Now, nums = [0,4,0,2,4].
// In the second operation, choose x = 2. Now, nums = [0,2,0,0,2].
// In the third operation, choose x = 2. Now, nums = [0,0,0,0,0].
	fmt.Println("Result := ",minimumOperations(nums))
}


func minimumOperations(nums []int) int {
    temp := make(map[int]struct{})
    for i := 0 ; i < len(nums) ; i++ {
        if _,ex := temp[nums[i]]; nums[i] != 0 && !ex{
            temp[nums[i]] = struct{}{}
        }
    }
    return len(temp)
}