package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}

		prefix = prefix[:j]
		fmt.Printf("\niteration:%d\nprefix:%s\n",i,prefix)
		if prefix == "" {
			// No common prefix found, exit early
			break
		}
	}

	return prefix
}

func main() {
	// Example usage
	// input := []string{"flower","flow", "flight"}
	input := []string{"dracu","lengui", "danger"}
	result := longestCommonPrefix(input)
	fmt.Println("Longest Common Prefix:", result)
}
