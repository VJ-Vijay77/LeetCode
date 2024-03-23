package main

import "fmt"


func main(){
	fmt.Println("Result := ",fizzBuzz(5))
}

func fizzBuzz(n int) []string {
    var res []string
	for i:=1; i<=n; i++ {
		switch  {
		case i%3 == 0 && i%5 == 0:
			res = append(res, "FizzBuzz")
		case i%3 == 0:
			res = append(res, "Fizz")
		case i%5==0:
			res = append(res, "Buzz")
		default:
			res = append(res, fmt.Sprintf("%d",i))
		}
	}

	return res
}