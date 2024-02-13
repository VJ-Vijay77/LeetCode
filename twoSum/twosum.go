package main

import "fmt"

func main() {
	num := [5]int{4, 7, 1, 3, 2}
	t := 9
	a := test(num, t)
	fmt.Println(a)

}

func test(n [5]int, t int) []int {

	var a = []int{}
	for i := 0; i < 5; i++ {
		for j := i+1 ; j <5 ; j++{
			
			if n[i] + n[j] == t {
				a = []int{i,j}
				break
			}
		}
	}

	return a
}
