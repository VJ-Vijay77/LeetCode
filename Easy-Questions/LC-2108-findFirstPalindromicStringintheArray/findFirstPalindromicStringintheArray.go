package main

import "fmt"

func main() {
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	// Output: "ada"
	// Explanation: The first string that is palindromic is "ada".
	// Note that "racecar" is also palindromic, but it is not the first.
	fmt.Println("Result := ", firstPalindrome(words))

}

func firstPalindrome(words []string) string {
	var ans string
	var yesNo bool

	for _, v := range words {
		if len(v) == 1 {
			return v
		}
		left, right := 0, len(v)-1
		for left < right {
			yesNo = true
			if v[left] != v[right] {
				yesNo = false
				break
			}
			left++
			right--
		}
		if !yesNo {
			continue
		} else {
			ans = v
			break
		}
	}
	return ans
}
