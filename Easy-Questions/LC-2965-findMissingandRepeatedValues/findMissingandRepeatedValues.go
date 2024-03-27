package main

import "fmt"

func main() {
	grid := [][]int{{9,1,7},{8,9,2},{3,4,6}}
	// Output: [2,4]
	// Explanation: Number 2 is repeated and number 4 is missing so the answer is [2,4].
	fmt.Println("Result := ",findMissingAndRepeatedValues(grid))
}


func findMissingAndRepeatedValues(grid [][]int) []int {
    max,min := 0,99999
	var res []int
	duplicate := make(map[int]int)
	for _,v := range grid{
		for _,k := range v{
			duplicate[k]++
			if k < min {
				min = k
			}
			if k > max {
				max = k
			}
		}
	}
	for i,v := range duplicate{
		if v > 1 {
			res = append(res, i)
			duplicate[i]--
		}
	}
	var sum int
	for i,_ := range duplicate{
		sum += i
	}
	if max*(max+1)/2 == sum {
		res = append(res, max+1)
	}else{
		res = append(res, max*(max+1)/2-sum)
	}
	return res
}