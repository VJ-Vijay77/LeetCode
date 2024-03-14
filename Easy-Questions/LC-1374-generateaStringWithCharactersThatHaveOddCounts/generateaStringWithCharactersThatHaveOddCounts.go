package main

import "fmt"

func main() {
	 n := 4
	// Output: "pppz"
	// Explanation: "pppz" is a valid string since the character 'p' occurs three times and the character 'z' occurs once. Note that there are many other valid strings such as "ohhh" and "love".
	fmt.Println(generateTheString(n))
}

func generateTheString(n int) string {
var res string

if n%2 == 0{
	for i:=0; i<n-1; i++{
		res += "a"
	}
	res += "b"
}else{
	for i:=0;i<n;i++{
		res += "a"
	}
}
return res
}
