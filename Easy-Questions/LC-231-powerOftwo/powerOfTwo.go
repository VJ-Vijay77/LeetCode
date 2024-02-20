package main

import (
	"fmt"
)

func main() {
	n := 12
	// Output: true
	// Explanation: 20 = 1
	fmt.Println("Result := ", isPowerOfTwo(n))

}

func isPowerOfTwo(n int) bool {
	base := 2
	if n == 1 || n==2{
		return true
	}
	for i := 1; i <31; i++ {
		
		for j:=0; j<i; j++ {
			
			base *= 2
		if base == n {
			return true
		}
		}		
	}

	return false
}


// more optimised
func isPowerOfTwo1(n int) bool {
	if n<=0{
		return false 
	}  
	  if n==1{
		  return true
	  }
	  if n%2==1{
		  return false
	  }
	  n=n/2
	  res:= isPowerOfTwo(n)
	  return res
  }