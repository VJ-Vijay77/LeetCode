package main

import "fmt"

func main() {
	encoded := []int{1,2,3}
	first := 1
	// Output: [1,0,2,1]
	// Explanation: If arr = [1,0,2,1], then first = 1 and encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]
	fmt.Println("Result := ",decode(encoded,first))
}

func decode(encoded []int, first int) []int {
	encoded = append([]int{first},encoded... )
	for i:=0; i<len(encoded)-1; i++{
		encoded[i+1] = encoded[i+1]^encoded[i]
	}
	return encoded
}