package main

import "fmt"


func main() {
	num := []int{1,2,1}
	//expected result : num[1,2,1,1,2,1]
	result := getConcatenation(num)
	fmt.Println(result)
}

// solution one, less optimised
func concatenateArray(num []int)[]int{
	n := 2*len(num)
	var res = make([]int,n)
	j:=0
	for i:=0; i<n; i++ {
		if j == len(num) {
			j = 0
		}
		res[i] = num[j] 
		j++
	}

	return res
}

// solution two, more optimized
func getConcatenation(num []int) []int {
    n := 2 * len(num)
    res := make([]int, n)
    for i := range res {
        res[i] = num[i%len(num)]
    }
    return res
}
