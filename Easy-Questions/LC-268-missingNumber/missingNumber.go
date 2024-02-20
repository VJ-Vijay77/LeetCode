package main

import "fmt"

func main() {
nums := []int{9,6,4,2,3,5,7,0,1}
// Output: 8
// Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.

	fmt.Println("Result := ",missingNumber(nums))
}


func missingNumber(nums []int) int {
    var ans int
	var yesNo bool

		for j:=1; j<=len(nums); j++{
			yesNo = false
			for i:=0; i<len(nums); i++{
				if j == nums[i] {
					yesNo = true
				} 
			}
			if !yesNo{
				ans = j
				break
			}
		}
	

	return ans
}

//more optimised
func missingNumber1(nums []int) int {
	n := len(nums)
	expected_sum := (n * (n + 1)) / 2
	calculated_sum := 0

	for i := 0; i < n; i++ {
		calculated_sum += nums[i]
	}
	return expected_sum - calculated_sum

}