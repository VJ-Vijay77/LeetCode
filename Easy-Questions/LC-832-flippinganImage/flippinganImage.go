package main

import "fmt"

func main() {
	image := [][]int{{1,1,0},{1,0,1},{0,0,0}}
	// Output: [[1,0,0],[0,1,0],[1,1,1]]
	// Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
	// Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
	fmt.Println("Result := ",flipAndInvertImage(image))
}

func flipAndInvertImage(image [][]int) [][]int {
	res := make([][]int,len(image))

	for i,v := range image{
		for j:=len(v)-1; j>=0; j--{
			temp := 0
			if v[j] == 0{
				temp = 1
			}
			res[i] = append(res[i], temp)
		}
	}

	return res
}
