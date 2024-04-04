package main

import "fmt"


func main() {
	arr := []string{"d","b","c","b","c","a"}
	 k := 2
	// Output: "a"
	// Explanation:
	// The only distinct strings in arr are "d" and "a".
	// "d" appears 1st, so it is the 1st distinct string.
	// "a" appears 2nd, so it is the 2nd distinct string.
	// Since k == 2, "a" is returned. 
	fmt.Println("Result := ",kthDistinct(arr,k))
}



func kthDistinct(arr []string, k int) string {
    res, strFrequency := "", make(map[string]int)
    
    for i := 0; i < len(arr); i++ {
        strFrequency[arr[i]]++
    }
    
    for i := 0; i < len(arr); i++ {
        if strFrequency[arr[i]] == 1 {
            k--
        }
        
        if k == 0 {
            res = arr[i]
            break
        }
    }
    
    return res
}