package main

import "fmt"

func main() {
	nums := []int{5, 12, 53, 22, 7, 59, 41, 54, 71, 24, 91, 74, 62, 47, 20, 14, 73, 11, 82, 2, 15, 38, 38, 20, 57, 70, 86, 93, 38, 75, 94, 19, 36, 40, 28, 6, 35, 86, 38, 94, 4, 90, 18, 87, 24, 22}
	// Output: [3,1]
	// Explanation:
	// Form a pair with nums[0] and nums[3] and remove them from nums. Now, nums = [3,2,3,2,2].
	// Form a pair with nums[0] and nums[2] and remove them from nums. Now, nums = [2,2,2].
	// Form a pair with nums[0] and nums[1] and remove them from nums. Now, nums = [2].
	// No more pairs can be formed. A total of 3 pairs have been formed, and there is 1 number leftover in nums.
	fmt.Println("Result := ", numberOfPairs(nums))
}

func numberOfPairs(nums []int) []int {
	var res []int
	pairCount := make(map[int]int)
	for _, v := range nums {
		pairCount[v]++
	}
	count := 0

	for i, v := range pairCount {
		if v == 2 {
			count++
			delete(pairCount, i)
		}
		if v > 2 {
			for pairCount[i] >= 2 {
				count++
				pairCount[i] -= 2
			}
			if pairCount[i] == 0 {
				delete(pairCount, i)
			}
		}
	}
	res = append(res, count, len(pairCount))
	return res
}
